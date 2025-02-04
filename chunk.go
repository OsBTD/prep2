package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	var result [][]int
	if size == 0 {
		fmt.Println()
	} else {
		for i := 0; i < len(slice); i += size {
			var x int
			x = i + size

			if x < len(slice) {
				result = append(result, slice[i:x])

			} else {
				result = append(result, slice[i:])
			}

		}
		fmt.Println(result)

	}

}
