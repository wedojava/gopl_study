// Generic archive file-reading function capable of reading ZIP files and tar files.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"gopl.io/ch10/exercise/ex10.2/archive"
)

func main() {
	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		headers, err := archive.List(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(sprintFileHeaders(headers))
	}
}

func sprintFileHeaders(headers []archive.FileHeader) string {
	var b bytes.Buffer
	namelen, sizelen := longestLength(headers)
	for _, header := range headers {
		fmt.Fprintf(&b, "% -*s %*d\n", namelen, header.Name, sizelen, header.Size)
	}
	return b.String()
}

func longestLength(headers []archive.FileHeader) (name int, size int) {
	for _, header := range headers {
		n := utf8.RuneCountInString(header.Name)
		if name < n {
			name = n
		}
		s := len(fmt.Sprintf("%d", header.Size))
		if size < s {
			size = s
		}
	}
	return
}
