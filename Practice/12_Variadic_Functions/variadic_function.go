package main

import "fmt"

func sum(nums ...int) int {
	sm := 0
	for _, num := range nums {
		sm += num
	}
	return sm
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
