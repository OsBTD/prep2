package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		return
	}
	for i := 0; i < len(args[1]); i++ {
		if args[1][i] < '0' || args[1][i] > '9' {
			return
		}
	}
	x := Atoi(args[1])
	if x <= 1 {
		return
	}
	var s string
	if isPrime(x) {
		fmt.Println(x)
		return
	} else {
		for i := 2; i <= x; i++ {
			if isPrime(i) {
				for x%i == 0 {
					s += strconv.Itoa(i) + "*"
					x = x / i
				}
			}

		}

	}
	fmt.Print(s[:len(s)-1])

}

func Atoi(s string) int {
	var v, conv int
	sign := 1
	if s[0] == '-' {
		s = s[1:]
		sign = -1
	}

	for _, c := range s {
		v = int(c - 48)
		conv += v
		conv = conv * 10
	}
	return conv / 10 * sign
}
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
