package app

import "strings"

func notBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}
