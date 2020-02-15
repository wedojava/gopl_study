package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// Ctrl + d to end the input
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			d := 0
			for _, n := range counts {
				if n > 1 {
					d++
				}
			}
			if len(counts) > 0 {
				fmt.Printf("file: %s have %d Duplicated lines.\n", arg, d)
			}
			_ = f.Close()
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
