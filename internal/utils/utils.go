package utils

// ToPtr takes a value of any type and returns a pointer to that value.
func ToPtr[T any](value T) *T {
	return &value
}
