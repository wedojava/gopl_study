package tar

import (
	"archive/tar"
	"io"
	"os"

	"gopl.io/ch10/exercise/ex10.2/archive"
)

func list(f *os.File) ([]archive.FileHeader, error) {
	var headers []archive.FileHeader

	// Open the tar archive for reading
	tr := tar.NewReader(f)
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return nil, err
		}
		headers = append(headers, archive.FileHeader{
			Name: hdr.Name,
			Size: uint64(hdr.Size),
		})
	}
	return headers, nil
}

// archive.formats will append tar format after blank import tar
// so, just a blank import make archive be able to recognized the file format after this init func.
func init() {
	archive.RegisterFormat("tar", "ustar\x0000", 257, list)
}
