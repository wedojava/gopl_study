package main

import (
	"fmt"
)

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(squares()())
	fmt.Println(squares()())
	fmt.Println(squares()())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
