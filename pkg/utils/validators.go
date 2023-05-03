package utils

import (
	"regexp"
	"unicode"
)

func IsEmptyOrWhitespace(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func IsEmail(s string) bool {
	// This is a simple email validation regex pattern that matches most common email formats
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, s)
	return match && err == nil
}
