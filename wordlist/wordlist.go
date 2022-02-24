package wordlist

import (
	"bufio"
	"embed"
	"strings"

	"github.com/pkg/errors"
)

//go:embed words_en.txt words_de.txt
var efs embed.FS

type WordList struct {
	words map[string]struct{}
}

func NewEN() (*WordList, error) {
	return newWordList("words_en.txt")
}

func NewDE() (*WordList, error) {
	return newWordList("words_de.txt")
}

func (wl *WordList) List(size int) []string {
	var ls []string
	for k := range wl.words {
		w := []rune(strings.TrimSpace(k))
		if len(w) == size {
			ls = append(ls, string(w))
		}
	}
	return ls
}

func newWordList(file string) (*WordList, error) {
	f, err := efs.Open(file)
	if err != nil {
		return nil, errors.Wrapf(err, "open embedded %q", file)
	}
	defer f.Close()

	wl := &WordList{
		words: map[string]struct{}{},
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \r\n\t")
		if s == "" {
			continue
		}
		wl.words[s] = struct{}{}
	}
	return wl, nil
}

func (wl *WordList) Contains(word string) bool {
	_, ok := wl.words[word]
	return ok
}

func (wl *WordList) Random() string {
	for k := range wl.words {
		return k
	}
	return ""
}

func (wl *WordList) RandomOfSize(size int) string {
	for k := range wl.words {
		if len(k) == size {
			return k
		}
	}
	return ""
}

func matches(w string, pattern string) bool {
	if len(w) != len(pattern) {
		return false
	}
	for i, b := range []byte(w) {
		pb := pattern[i]
		if pb == '*' {
			continue
		}
		if pb != b {
			return false
		}
	}
	return true
}

func (wl *WordList) Find(pattern string) []string {
	var ms []string
	for k := range wl.words {
		if matches(k, pattern) {
			ms = append(ms, k)
		}
	}
	return ms
}

func (wl *WordList) AllOfSize(size int) []string {
	var ms []string
	for k := range wl.words {
		if len([]rune(k)) == size {
			ms = append(ms, k)
		}
	}
	return ms
}
