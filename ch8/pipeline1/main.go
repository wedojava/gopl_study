package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()
	//Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
