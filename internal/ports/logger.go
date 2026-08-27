package ports

// Logger defines the interface for logging messages at different levels (Debug, Info, Warn, Error).
type Logger interface {
	// Debug logs a message at the debug level.
	Debug(msg string, fields ...any)
	// Info logs a message at the info level.
	Info(msg string, fields ...any)
	// Warn logs a message at the warn level.
	Warn(msg string, fields ...any)
	// Error logs a message at the error level.
	Error(msg string, fields ...any)
}
