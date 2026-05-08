package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func allLangs() (string, string, string) {
	return "GoLang", "Python", "C++"
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero.")
	}

	return a / b, nil
}

func solution(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	fmt.Println(add(2, 3))

	lang1, lang2, lang3 := allLangs()
	fmt.Println(allLangs())
	fmt.Println(lang1, lang2, lang3)

	add := func(a, b int) int {
		return a + b
	}
	subtract := func(a, b int) int {
		return a - b
	}
	multiply := func(a, b int) int {
		return a * b
	}
	// divide := func(a, b int) (int, error) {
	// 	if b == 0 {
	// 		return 0, fmt.Errorf("Can't divide by zero")
	// 	}
	// 	return a / b, nil
	// }

	fmt.Println(divide(4, 0))

	a, b := 4, 5
	fmt.Println(solution(a, b, add))
	fmt.Println(solution(a, b, subtract))
	fmt.Println(solution(a, b, multiply))
	// fmt.Println(solution(a, b, divide))

}
