package main

import "fmt"

// arrays are memory efficient, fixed size, contact time access

func main() {
	// var arr [4]int
	// arr[1] = 5
	// fmt.Println(arr)

	arr := [4]int{}
	fmt.Println(len(arr))

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2d)
}
