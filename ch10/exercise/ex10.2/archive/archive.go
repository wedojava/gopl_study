package archive

import (
	"os"
)

type FileHeader struct {
	Name string
	Size uint64
}

func List(f *os.File) ([]FileHeader, error) {
	format, err := sniff(f)
	if err != nil {
		return nil, err
	}
	return format.list(f)
}
