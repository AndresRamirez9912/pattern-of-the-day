package internal

var (
	// Version is the version of the project registered when a new
	// tag is created
	Version = "unknown"

	// Commit is the message of the last commit or tag creation
	Commit = "unknown"

	// BuildDate is the date when the application was built
	BuildDate = "unknown"
)

// VersionInfo models the basic information about the project development
type VersionInfo struct {
	Version   string
	COmmit    string
	BuildDate string
}
