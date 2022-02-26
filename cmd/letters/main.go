package main

import (
	"fmt"
	"sort"

	"github.com/mazzegi/worpler"
)

func main() {
	ls := rankLetters(worpler.WordlistAnswers5)
	fmt.Printf("EN: %q\n", string(ls))
	ls = rankLetters(worpler.WordlistAnswers5_DE)
	fmt.Printf("DE: %q\n", string(ls))
}

func rankLetters(list []string) []rune {
	h := map[rune]int{}
	for _, w := range list {
		for _, r := range w {
			h[r]++
		}
	}
	type rn struct {
		r rune
		n int
	}
	var sl []rn
	var total int
	for r, n := range h {
		sl = append(sl, rn{r, n})
		total += n
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].n > sl[j].n
	})

	var ls []rune
	for _, e := range sl {
		ls = append(ls, e.r)
		fmt.Printf("%q => %f\n", string(e.r), float64(e.n)/float64(total))
	}
	return ls
}
