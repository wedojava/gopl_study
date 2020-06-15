package limitreader

import (
	"io"
)

type MyLimitReader struct {
	Reader io.Reader
	Limit  int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r, n}
}

func (m *MyLimitReader) Read(p []byte) (n int, err error) {
	if m.Limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > m.Limit {
		p = p[:m.Limit]
	}
	n, err = m.Reader.Read(p)
	m.Limit -= int64(n)
	return
}
