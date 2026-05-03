package main

import "fmt"

func main() {
	// go doesn't have while loop
	i := 1

	//standard for
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	//while loop
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//infinite while loop
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

}
