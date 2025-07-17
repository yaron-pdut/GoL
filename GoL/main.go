package main

import (
	"fmt"
	"sync"
)

func printMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from goroutine!")
}

func main() {
	fmt.Println("Hello, World!")
	var wg sync.WaitGroup
	wg.Add(1)
	go printMessage(&wg)
	wg.Wait()
}
