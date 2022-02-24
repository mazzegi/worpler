package worpler

import (
	"strings"
	"unicode"

	"github.com/mazzegi/worpler/wordlist"
)

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
	for _, word := range wordlist5 {
		if match(strings.ToLower(word), strings.ToLower(pattern), strings.ToLower(exclude), strings.ToLower(include)) {
			matches = append(matches, word)
		}
	}
	return matches
}

//
func matchV2(word string, pattern string, exclude string) bool {
	word = strings.ToLower(word)
	exclude = strings.ToLower(exclude)

	rsword := []rune(word)
	rswordIncludeCheck := []rune(word)
	rspattern := []rune(pattern)
	if len(rsword) != len(rspattern) {
		return false
	}
	if strings.ContainsAny(word, exclude) {
		return false
	}

	indexOf := func(rs []rune, fr rune) int {
		for i, r := range rs {
			if r == fr {
				return i
			}
		}
		return -1
	}

	// first match exact positions
	for i, r := range pattern {
		if r == '*' {
			// wildcard * matches any
			continue
		}
		matchPos := unicode.IsUpper(r)
		if !matchPos {
			continue
		}
		if rsword[i] != unicode.ToLower(r) {
			return false
		}
		// mark with uppercase char
		rswordIncludeCheck[i] = r
	}

	// last match deviant positions
	for i, r := range pattern {
		if r == '*' {
			// wildcard * matches any
			continue
		}
		matchPos := unicode.IsUpper(r)
		if matchPos {
			continue
		}
		mi := indexOf(rswordIncludeCheck, r)
		switch {
		case mi < 0:
			return false
		case mi == i:
			// found on same pos -> not match
			return false
		default:
			rswordIncludeCheck[mi] = unicode.ToUpper(r)
		}
	}
	return true
}

// Find finds all words matching pattern, exluding those, which contain chars of "exclude"
// Uppercase letters in pattern means, that they are already at their designation position
// Lowercase letters in pattern means, that they occur in the answer, but on a different position
func FindV2(list []string, pattern string, exclude string) []string {
	var matches []string
	//for _, word := range wordlistAnswers5 {
	for _, word := range list {
		if matchV2(word, pattern, exclude) {
			matches = append(matches, word)
		}
	}
	return matches
}

//
func matchSpebee(word string, letters string) bool {
	rword := []rune(letters)
	for _, rl := range word {
		ir := strings.IndexRune(string(letters), rl)
		if ir < 0 {
			return false
		}
		rword[ir] = '_'
	}
	return true
}

//
func FindSpebee(letters string) []string {
	wl, _ := wordlist.NewEN()
	letters = strings.ToLower(letters)
	var matches []string
	for n := 4; n < len(letters); n++ {
		words := wl.AllOfSize(n)
		for _, word := range words {
			if matchSpebee(strings.ToLower(word), letters) {
				matches = append(matches, word)
			}
		}
	}
	return matches
}
