package main

import (
	"fmt"
	"slices"
)

// slices package
// slice -> dynamic, useful in go
func main() {
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil) // NULL in go is nil

	new_nums := [][]int{{1, 2, 3}, {4, 5}}
	fmt.Println(new_nums)

	var a = make([]int, 2, 5)

	// fmt.Println(a)
	// fmt.Println(cap(a))
	// fmt.Println(len(a))

	// a = append(a, 5)
	// a = append(a, 6)
	// a = append(a, 8)
	// a = append(a, 7)

	// fmt.Println(a[2:])
	// fmt.Println(cap(a))
	// fmt.Println(len(a))

	arr := []int{1, 2, 3} //auto size
	fmt.Println(arr)

	for index, value := range arr {
		fmt.Println(index, value)
	}

	fmt.Println(slices.Equal(arr, a))

}
