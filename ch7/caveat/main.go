// page 185 Caveat: An Interface Containing a Nil Pointer Is Non-Nil
package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	// var buf io.Write
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
		fmt.Println(out)
	}
}
