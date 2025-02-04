package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 2 || Atoi(args[1]) < 0 {
		fmt.Println(0)
	} else {
		x := Atoi(args[1])
		fmt.Println(Sum(x))

	}

}
func Sum(x int) int {
	var sum int
	for i := x; i > 0; i-- {
		if isPrime(i) {
			sum += i
		}
	}
	return sum

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
