package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	for ind, val := range nums {
		fmt.Println(ind, val)
	}

	mp := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println(mp)

	for key, value := range mp {
		fmt.Println(key, value)
	}

	for i, c := range "Aditi" {
		fmt.Println(i, string(c))
	}
}
