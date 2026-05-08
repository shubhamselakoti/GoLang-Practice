package main

import "fmt"

func main() {

	// i := 2
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// tmp := func(i interface{}) {
	// 	switch i.(type) {
	// 	case string:
	// 		fmt.Println("String")
	// 	case int:
	// 		fmt.Println("int")
	// 	default:
	// 		fmt.Println("others")
	// 	}
	// }
	// tmp("hello")

	day := 0

	switch day {
	case 0, 6:
		fmt.Println("Weekend")
	default:
		fmt.Println("weekday")
	}
}
