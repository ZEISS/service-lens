package utils

import (
	"strconv"

	htmx "github.com/zeiss/fiber-htmx"
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

// Resolvers ...
func Resolvers(funcs ...htmx.ResolveFunc) htmx.Config {
	return htmx.Config{
		Resolvers: funcs,
	}
}
