package utils

import (
	"strconv"
)

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

// IntStr converts the given int to a string.
func IntStr(i int) string {
	return strconv.Itoa(i)
}

// StrInt converts the given string to an int.
func StrInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
