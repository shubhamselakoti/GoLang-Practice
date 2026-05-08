package main

import "fmt"

// func printTheSliceInt(nums []int) {
// 	for ind, val := range nums {
// 		fmt.Println(ind, val)
// 	}
// }

// func printTheSliceString(names []string) {
// 	for ind, val := range names {
// 		fmt.Println(ind, val)
// 	}
// }

func printTheSlice[T any](slice []T) { //now we don't have to make different different func to accept the datatypes
	for ind, val := range slice {
		fmt.Println(ind, val)
	}
}

// func printTheSlice[T int | string](slice []T) { // liniting the datatype entry
// 	for ind, val := range slice {
// 		fmt.Println(ind, val)
// 	}
// }

func main() {
	nums := []int{1, 2, 3, 4}
	names := []string{"a", "b", "c", "d"}

	// printTheSliceInt(nums)
	// printTheSliceString(names)
	printTheSlice(nums)
	printTheSlice(names)
}
