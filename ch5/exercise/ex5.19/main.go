// start
// Now defer start
// oops: 1
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "oops: %d\n", oops())
}

func oops() (r int) {
	fmt.Println("start")
	defer func() {
		fmt.Println("Now defer start")
		if p := recover(); p != nil {
			r = 1
		}
	}()
	panic("oops!")
}
