// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	//s, sep := "", ""

	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
		//s += sep + arg
		//sep = " "
	}
	//fmt.Println(s)
}
