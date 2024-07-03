package main

import "strings"

/*
Instructions
Your task is to determine what Bob will reply to someone when they say something to him or ask him a question.

Bob only ever answers one of five things:

"Sure." This is his response if you ask him a question, such as "How are you?" The convention used for questions is that it ends with a question mark.
"Whoa, chill out!" This is his answer if you YELL AT HIM. The convention used for yelling is ALL CAPITAL LETTERS.
"Calm down, I know what I'm doing!" This is what he says if you yell a question at him.
"Fine. Be that way!" This is how he responds to silence. The convention used for silence is nothing, or various combinations of whitespace characters.
"Whatever." This is what he answers to anything else.
*/
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
