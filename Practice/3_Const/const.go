package main

import "fmt"

func main() {
	const name = "Shubham" // reasign not allowed!!!

	const (
		age  = 22
		good = true
	)

	fmt.Println(name, age, good)
}
