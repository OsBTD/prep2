package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	// input2 := []uint{3, 2, 1, 0, 4}
	// fmt.Println(CanJump(input2))

	// input3 := []uint{0}
	// fmt.Println(CanJump(input3))

	input4 := []uint{1, 1, 1, 1, 0}
	input5 := []uint{5, 4, 3, 2, 1, 0}
	input6 := []uint{0}
	input7 := []uint{5}
	input8 := []uint{}
	input9 := []uint{1, 2, 3, 0, 2}
	input10 := []uint{3, 2, 1, 0, 4}
	input11 := []uint{0, 0, 0, 0, 0}
	input12 := []uint{1, 2, 3}
	input13 := []uint{1, 2, 3, 0, 1}
	input14 := []uint{1, 0, 0, 0, 0}
	input15 := []uint{1, 0, 1, 0, 1}
	input16 := []uint{10, 20, 30, 40, 0}
	fmt.Println(CanJump(input4))
	fmt.Println(CanJump(input5))
	fmt.Println(CanJump(input6))
	fmt.Println(CanJump(input7))
	fmt.Println(CanJump(input8))
	fmt.Println(CanJump(input9))
	fmt.Println(CanJump(input10))
	fmt.Println(CanJump(input11))
	fmt.Println(CanJump(input12))
	fmt.Println(CanJump(input13))
	fmt.Println(CanJump(input14))
	fmt.Println(CanJump(input15))
	fmt.Println(CanJump(input16))

}
func CanJump(c []uint) bool {
	if len(c) == 0 {
		return false
	}
	i := 0
	for i < len(c)-1 {
		if c[i] == 0 {
			return false
		}
		if i+int(c[i]) > len(c)-1 {
			return false
		} else {
			i += int(c[i])
		}

	}
	return true
}
