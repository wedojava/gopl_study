package main

import (
	"fmt"
	"os"

	"gopl.io/ch4/exercise/ex4.13/omdb"
)

func main() {
	if err := omdb.GetPoster(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "ch04/exercise/ex4.13: %v\n", err)
	}
}
