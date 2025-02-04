package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	for i := 1; i < len(args); i++ {
		fmt.Println(RevstrCap(args[i]))
	}
}

func isSpace(c rune) bool {
	return c == ' '
}

func isCap(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func isLow(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func Upper(c rune) rune {
	if isLow(c) {
		c = c - 32
	}
	return c
}

func Lower(c rune) rune {
	if isCap(c) {
		c = c + 32
	}
	return c
}

func split(s string) []string {
	result := ""
	var res []string
	for i := 0; i < len(s); i++ {
		if isSpace(rune(s[i])) {
			if result != "" {
				res = append(res, result)
				result = ""
			}
		} else {
			result += string(s[i])
		}
	}
	if result != "" {
		res = append(res, result)
	}

	return res
}

func RevstrCap(s string) string {
	words := split(s)
	var res string
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j != len(words[i])-1 {
				if isCap(rune(words[i][j])) {
					res += string(Lower(rune(words[i][j])))
				} else {
					res += string(words[i][j])
				}
			} else {
				if i != len(words)-1 {
					if isLow(rune(words[i][j])) {
						res += string(Upper(rune(words[i][j]))) + " "
					} else {
						res += string(words[i][j]) + " "
					}
				} else {
					if isLow(rune(words[i][j])) {
						res += string(Upper(rune(words[i][j])))
					} else {
						res += string(words[i][j])
					}
				}
			}
		}
	}
	return res
}
