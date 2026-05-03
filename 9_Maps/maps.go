package main

import (
	"fmt"
	"maps"
)

// maps package
// map ->  hash

func main() {

	m := make(map[int]string) // map[key]value
	mp := map[string]string{} // another way
	fmt.Println(maps.Equal(mp, mp))
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"

	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}

	key := 1
	val, ok := m[key]

	if ok {
		fmt.Println(key, "exist : ", val)
	}

	delete(m, 3)

	for key := range m {
		fmt.Println(m[key])
	}

	clear(m)

	for key := range m {
		fmt.Println(m[key])
	}
}
