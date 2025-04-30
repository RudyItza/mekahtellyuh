package app

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

// notBlank checks if a string is not empty after trimming spaces.
func notBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

// ValidateInputs checks form fields and returns a map of validation errors.
func ValidateInputs(title, content, username, email string) map[string]string {
	errors := make(map[string]string)

	// Title validation
	if !notBlank(title) {
		errors["title"] = "Title cannot be blank."
	} else if len(title) < 5 || len(title) > 30 {
		errors["title"] = "Title must be between 5 and 30 characters."
	}

	// Content validation
	if !notBlank(content) {
		errors["content"] = "Content cannot be blank."
	} else if len(content) > 800 {
		errors["content"] = "Content must be at least 800 characters."
	}

	// Username validation
	if !notBlank(username) {
		errors["username"] = "Username cannot be blank."
	} else {
		if !usernameRegex.MatchString(username) {
			errors["username"] = "Username must contain only letters, numbers, or underscores."
		} else if len(username) < 3 || len(username) > 20 {
			errors["username"] = "Username must be between 3 and 20 characters."
		}
	}

	// Email validation
	if !notBlank(email) {
		errors["email"] = "Email cannot be blank."
	} else if !emailRegex.MatchString(email) {
		errors["email"] = "Invalid email address."
	}

	return errors
}
