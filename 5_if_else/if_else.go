package main

import "fmt"

func main() {
	// go doen't have ternary operator
	age := 18
	if age >= 18 {
		fmt.Println("yes")
	} else if age >= 12 {
		fmt.Println("Not yet")
	} else {
		fmt.Println("kiddo")
	}

	if num := 7; num == 7 {
		fmt.Println("good")
	} else {
		fmt.Println("not good")
	}
}
