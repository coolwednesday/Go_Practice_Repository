package main

import "strings"

// Hey should have a comment documenting it.
func Hey(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Fine. Be that way!"
	} else if strings.ToUpper(s) == s && strings.ToLower(s) != strings.ToUpper(s) {
		if strings.Index(s, "?") == len(s)-1 {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	} else if strings.Index(s, "?") == len(s)-1 {
		return "Sure."
	}
	return "Whatever."
}
