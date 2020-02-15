package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// Ctrl + d to end the input
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		// `counts[input.Text()]++` is equivalent to these two statements:
		// line := input.Text()
		// counts[line] = counts[line] + 1
		if input.Text() == "EOF" {
			break
		}
	}
	//	Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
