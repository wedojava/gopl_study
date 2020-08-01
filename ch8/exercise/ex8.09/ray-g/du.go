package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type RootSize struct {
	root int
	size int64
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Traverse each root of the file tree in parallel.
	rootSizes := make(chan RootSize)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, i, rootSizes)
	}
	go func() {
		n.Wait()
		close(rootSizes)
	}()

	// Print the results periodically.
	// tick := time.Tick(500 * time.Millisecond)
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutine to finish.
			for range rootSizes {
				// Do Nothing
			}
			return
		case rs, ok := <-rootSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[rs.root]++
			nbytes[rs.root] += rs.size
			// case <-tick:
			//         printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%d files %.1f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

func walkDir(dir string, n *sync.WaitGroup, root int, rootSizes chan<- RootSize) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, root, rootSizes)
		} else {
			rootSizes <- RootSize{root, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
