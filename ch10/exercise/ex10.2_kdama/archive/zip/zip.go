package zip

import (
	"archive/zip"
	"os"

	"gopl.io/ch10/exercise/ex10.2_kdama/archive"
)

func list(f *os.File) ([]archive.FileHeader, error) {
	var headers []archive.FileHeader

	// Open a zip archive for reading
	r, err := zip.OpenReader(f.Name())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		headers = append(headers, archive.FileHeader{
			Name: f.Name,
			Size: f.UncompressedSize64,
		})
	}
	return headers, nil
}

func init() {
	archive.RegisterFormat("zip", "PK\x03\x04", 0, list)
	archive.RegisterFormat("zip", "PK\x05\x06", 0, list)
}