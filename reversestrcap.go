package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
}

// incomplete
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
		c = c + 32
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
