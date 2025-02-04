package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5}))

	// 	[1 4 2 5 3 6]
	// [1 2 3 4 5 6 7 8 9 10 11]
	// [4 1 5 2 6 3 7 8 9]
	// [1 2 3]

}

func ConcatAlternate(slice1, slice2 []int) []int {
	var s bool
	var result []int
	if len(slice1) == 0 {
		return slice2
	} else if len(slice2) == 0 {
		return slice1
	}
	maxlen := len(slice1)
	if len(slice2) > len(slice1) {
		maxlen = len(slice2)
		s = true
	}
	for i := 0; i < maxlen; i++ {
		if !s {

			result = append(result, slice1[i])
			if i < len(slice2) {
				result = append(result, slice2[i])

			}

		} else {
			result = append(result, slice2[i])
			if i < len(slice1) {
				result = append(result, slice1[i])

			}

		}

	}

	return result
}
