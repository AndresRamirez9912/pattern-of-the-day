package domain


// Supported errors
var (
	ErrMaxCluesReached = NewError("error max clues per challenge reached", 1)
)
