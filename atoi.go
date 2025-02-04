package main

import "fmt"

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
	return conv/10 * sign
}
func main() {
	a := "12345"
	b := "-123456"
	fmt.Println(Atoi(a))
	fmt.Println(Atoi(b))

}
