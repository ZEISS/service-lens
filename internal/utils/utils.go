package utils

// StrPtr returns a pointer to the given string.
func StrPtr(s string) *string {
	return &s
}

// PtrStr returns the value of the given string pointer.
func PtrStr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
