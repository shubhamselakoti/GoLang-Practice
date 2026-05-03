package main

import (
	"fmt"
	"sync"
)

// waitgoup cycle :: add -> done -> wait

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done() // wherever function exits it runs
	fmt.Println("Performing Task:", i)
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go task(i, &wg) // parallel execution // oncept of mutlithreading
	}
	wg.Wait()
}
