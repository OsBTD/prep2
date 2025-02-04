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
	var a string = args[1]
	var b string = args[2]
	fmt.Println(hiddenP(a, b))

}

func hiddenP(a, b string) int {
	if a == "" {
		return 1
	}
	i, j := 0, 0
	for j < len(b) {
		if b[j] == a[i] {
			i++
		}
		if i == len(a)-1 {
			return 1
		}
		j++
	}
	return 0
}
