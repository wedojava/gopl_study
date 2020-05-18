package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("index: %d \t arg: %v\n", i, arg)
	}
}
