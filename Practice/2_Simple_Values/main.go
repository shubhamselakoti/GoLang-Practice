package main

import "fmt"

var name string = "Shubham Selakoti"
var age = 22

//shorthand declaration
// isGood := true
func main() {
	printName()
	printAge()
	printGood()
}
func printName() {
	fmt.Println(name)
}
func printAge() {
	fmt.Println(age)
}
func printGood() {
	isGood := true
	fmt.Println(isGood)
}
