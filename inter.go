package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	a := args[1]
	b := args[2]
	fmt.Println(inter(a, b))
}

func inter(a, b string) string {
	var result string
	exists := make(map[rune]bool)
	for _, c := range b {
		exists[c] = true
	}
	found := make(map[rune]bool)
	for _, c := range a {
		if exists[c] && !found[c] {
			result += string(c)
			found[c] = true
		}
	}
	return result
}
