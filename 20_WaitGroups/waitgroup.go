package main

import (
	"fmt"
	"time"
)

func task(i int) {
	fmt.Println("Performing Task:", i)
}

func main() {
	for i := range 10 {
		go task(i) // parallel execution // oncept of mutlithreading
	}

	time.Sleep(time.Second)

}
