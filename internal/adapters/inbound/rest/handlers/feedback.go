package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/attempt"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/challenge"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/feedback"
)

// maxUploadMemory bounds how much of the multipart body is buffered in memory.
const maxUploadMemory = 32 << 20

// FeedbackHandler handles HTTP requests related to feedback.
type FeedbackHandler struct {
	createFeedback *feedback.CreateFeedbackUseCase
	getAttempt     *attempt.GetAttemptUseCase
	getChallenge   *challenge.GetChallengeUseCase
}

// NewFeedbackHandler creates a new instance of FeedbackHandler.
func NewFeedbackHandler(
	createFeedback *feedback.CreateFeedbackUseCase,
	getAttempt *attempt.GetAttemptUseCase,
	getChallenge *challenge.GetChallengeUseCase,
) *FeedbackHandler {
	return &FeedbackHandler{
		createFeedback: createFeedback,
		getAttempt:     getAttempt,
		getChallenge:   getChallenge,
	}
}

// Create handles a multipart upload: attempt_id as a form field, the
// solution as one or more files under the "files" field.
func (h *FeedbackHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request. This will store up to maxUploadMemory bytes
	// in memory and the rest on disk.
	err := r.ParseMultipartForm(maxUploadMemory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert attempt Id to an int64.
	attemptId, err := strconv.ParseInt(r.FormValue("attempt_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid attempt_id", http.StatusBadRequest)
		return
	}

	// Retrieve the uploaded files from the "files" form field.
	fileHeaders := r.MultipartForm.File["files"]
	if len(fileHeaders) == 0 {
		http.Error(w, "files is required", http.StatusBadRequest)
		return
	}

	// Uploaded files live on the server, unlike a client-side solution_path.
	solutionDir, err := writeUploadedFiles(fileHeaders)
	if err != nil {
		h.getAttempt.Logger.Error("error writing uploaded files", "error", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer os.RemoveAll(solutionDir)

	// Fetch attempt information from the provided attempt Id.
	foundAttempt, err := h.getAttempt.Execute(r.Context(), attemptId)
	if err != nil {
		h.getAttempt.Logger.Error("error executing get attempt use case", "error", err.Error())

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Ensure the attempt has an associated challenge.
	if foundAttempt.ChallengeId == nil {
		http.Error(w, "attempt has no associated challenge", http.StatusInternalServerError)
		return
	}

	// Execute the use case to get the challenge details.
	foundChallenge, err := h.getChallenge.Execute(r.Context(), *foundAttempt.ChallengeId)
	if err != nil {
		h.getChallenge.Logger.Error("error executing get challenge use case", "error", err.Error())

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Execute the use case to create the feedback.
	createdFeedback, err := h.createFeedback.Execute(r.Context(), foundAttempt, foundChallenge, solutionDir, ".")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.createFeedback.Logger.Info("feedback created successfully", "feedback_id", createdFeedback.Id)

	// Prepare the response payload with the created feedback and attempt details.
	resp := CreateFeedbackResponse{
		Feedback: Feedback{
			FeedbackId:  createdFeedback.Id,
			Score:       createdFeedback.Score,
			Summary:     createdFeedback.Summary,
			Suggestions: createdFeedback.Suggestions,
		},
		Attempt: Attempt{
			AttemptId: foundAttempt.Id,
			Status:    string(foundAttempt.Status),
			Sequence:  foundAttempt.SequenceOrder,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// writeUploadedFiles saves each uploaded file into a fresh temp directory.
func writeUploadedFiles(headers []*multipart.FileHeader) (string, error) {
	// Create a temporary directory to store the uploaded files.
	dir, err := os.MkdirTemp("", "solution-*")
	if err != nil {
		return "", err
	}

	// Iterate over each uploaded file header and write the file to the temporary directory.
	for _, header := range headers {
		err := writeUploadedFile(dir, header)
		if err != nil {
			os.RemoveAll(dir)
			return "", err
		}
	}

	return dir, nil
}

// writeUploadedFile copies one uploaded file into dir, rejecting names that would escape it.
func writeUploadedFile(dir string, header *multipart.FileHeader) error {
	cleaned := filepath.Clean(header.Filename)
	if filepath.IsAbs(cleaned) || cleaned == ".." || strings.HasPrefix(cleaned, ".."+string(filepath.Separator)) {
		return fmt.Errorf("invalid file name %q", header.Filename)
	}

	src, err := header.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Ensure the directory for the full path exists before creating the file.
	fullPath := filepath.Join(dir, cleaned)
	err = os.MkdirAll(filepath.Dir(fullPath), 0o755)
	if err != nil {
		return err
	}

	// Create the destination file for the uploaded content.
	dst, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy the contents of the uploaded file to the destination file.
	_, err = io.Copy(dst, src)
	return err
}
