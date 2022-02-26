package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mazzegi/worpler"
)

func main() {

	pattern := flag.String("p", "*****", "pattern to match (defaults to five wildcards)")
	exclude := flag.String("e", "", "chars to exclude")
	flag.Parse()

	fmt.Printf("finding all (5-letter) words matching %q, exluding %q\n", *pattern, *exclude)
	t0 := time.Now()
	//matches := worpler.Find(*pattern, *exclude, *include)
	matches := worpler.FindV2(worpler.WordlistAnswers5_DE, *pattern, *exclude)
	fmt.Printf("found %d matches in %s\n", len(matches), time.Since(t0))
	for _, m := range matches {
		fmt.Printf("%q\n", m)
	}
}
