package worpler

import (
	"sort"
)

var lettersRankEN = []rune("earotlisncuydhpmgbfkwvzxqj")

var lettersRankDE = []rune("eastrnliouhmkbgdpfczwäyüövjxßq")

func rank(r rune, lettersRank []rune) int {
	for i, lr := range lettersRank {
		if lr == r {
			return i + 1
		}
	}
	return len(lettersRank) + 2
}

func rankWord(w string, ranks []rune) int {
	var maxRepeat int
	h := map[rune]int{}
	var l int
	for _, r := range w {
		l += rank(r, ranks)
		h[r]++
		if h[r] > maxRepeat {
			maxRepeat = h[r]
		}
	}
	if maxRepeat > 0 {
		l = l * maxRepeat
	}
	return l
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
	}
	return ls
}
