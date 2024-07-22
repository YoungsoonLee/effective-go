package main

import (
	"fmt"
	"regexp"
	"unicode"
)

var (
	// Example: match "rN" is "userName"
	camelRe = regexp.MustCompile(`[a-z][A-Z]`)
)

// fixCase gets a string in the format "aB" and returns "a_b".
func fixCase(s string) string {
	return fmt.Sprintf("%c_%c", s[0], unicode.ToLower(rune(s[1])))
}

// camelToLower gets a camel case string and returns a snake case string.
func camelToLower(s string) string {
	return camelRe.ReplaceAllStringFunc(s, fixCase)
}

func main() {
	s := "userName"
	fmt.Println(camelToLower(s))
}
