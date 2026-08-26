package domain

import "fmt"

// Attempt defines the properties of an attempt
type Attempt struct {
	Id          int64
	FeedbackId  *int64
	ChallengeId *int64
	Status      AttemptStatus
}

// AttemptStatus is a custom type to define the supported attempt status
type AttemptStatus string

// Supported attempt status
const (
	AttemptStatusPending   AttemptStatus = "pending"
	AttemptStatusCompleted AttemptStatus = "completed"
	AttemptStatusFailed    AttemptStatus = "failed"
)

// Complete marks the attempt as completed
func (a *Attempt) Complete() error {
	// Validate if the attempt has a feedback associated
	if a.FeedbackId == nil {
		return fmt.Errorf("cannot complete attempt without feedback")
	}

	// Mark the attempt as completed
	a.Status = AttemptStatusCompleted

	return nil
}
