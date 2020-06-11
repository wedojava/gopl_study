// code cp from https://github.com/wedojava/gopl/blob/master/ch07/ex7.04/stringreader.go
// It's resemble to thought of https://golang.org/src/strings/reader.go?s=3567:3599#L144
package stringreader

import (
	"io"
)

type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &StringReader{s}
}
