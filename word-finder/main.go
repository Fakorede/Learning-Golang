package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "Lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}
