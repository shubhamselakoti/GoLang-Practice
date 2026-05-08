package main

import (
	"fmt"
	"my-application/auth"

	"github.com/fatih/color"
)

func get(name string) string {
	return "Hi " + name
}

func main() {
	fmt.Println("Hi!!!")
	res := auth.Login("aditikeshri12@gmail.com", "Choti@13")
	// a := auth.matchPass("Choti@123")

	// name := "Aditi"

	// b := get(name)

	// fmt.Println(b)
	// fmt.Println(res)
	if res == "Logged In" {
		color.Green(res)
	} else {
		color.Red("Invalid")
	}
}
