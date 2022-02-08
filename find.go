package worpler

import "strings"

func match(word string, pattern string, exclude string, include string) bool {
	rsword := []rune(word)
	rspattern := []rune(pattern)
	if len(rsword) != len(rspattern) {
		return false
	}
	if strings.ContainsAny(word, exclude) {
		return false
	}
	for _, r := range include {
		if !strings.ContainsRune(word, r) {
			return false
		}
	}
	for i, r := range pattern {
		if r == '*' {
			// wildcard * matches any
			continue
		}
		if rsword[i] != r {
			return false
		}
	}
	return true
}

// Find finds all words mathing pattern, exluding those, which contain chars of exclude, and exclude those which don't containschars of include.
// Find is case-insensitive.
func Find(pattern string, exclude string, include string) []string {
	var matches []string
	for _, word := range wordlist {
		if match(strings.ToLower(word), strings.ToLower(pattern), strings.ToLower(exclude), strings.ToLower(include)) {
			matches = append(matches, word)
		}
	}
	return matches
}
