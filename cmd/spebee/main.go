package main

import (
	"fmt"
	"os"

	"github.com/mazzegi/worpler"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: spebee <letters>", len(os.Args))
		return
	}

	ms := worpler.FindSpebee(os.Args[1])
	for _, m := range ms {
		fmt.Printf("%q\n", m)
	}
}
