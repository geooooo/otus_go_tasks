package main

import (
	"fmt"
)

// Написать функцию,
// которая получает несколько слайсов и склеивает их в один длинный.
//  { {1, 2, 3}, {4, 5}, {6, 7} } => {1, 2, 3, 4, 5, 6, 7}

func main() {
	slice := [][]int{{1, 2, 3}, {4, 5}, {6, 7}}
	fmt.Println(MakeFlatSlice(slice))
}

func MakeFlatSlice(slice [][]int) []int {
	var flatSlice []int

	for i := 0; i < len(slice); i++ {
		flatSlice = append(flatSlice, slice[i]...)
	}

	return flatSlice
}
