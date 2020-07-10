// go build gopl.io/ch3/mandelbrot > mandelbrot.bmp
// go run main.go -f gif < ./mandelbrot.bmp > out.gif
// go run main.go -f jpg < ./mandelbrot.bmp > out.jpg
// go run main.go -f png < ./mandelbrot.bmp > out.png
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder, so the code can recognize input in png format.
	"io"
	"log"
	"os"
)

func main() {
	var format string
	flag.StringVar(&format, "f", "", "output format, Required. One of png, jpg, gif.")
	flag.Parse()
	info, _ := os.Stdout.Stat()
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Fprintf(os.Stderr, "Refusing to write to character device. Redirect output to a pipe or regular file.")
		os.Exit(1)
	}
	if err := convert(os.Stdin, os.Stdout, format); err != nil {
		log.Fatalf("ch10/ex10.1: %v\n", err)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	switch format {
	case "gif":
		return gif.Encode(out, img, nil)
	case "jpg", "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
}
