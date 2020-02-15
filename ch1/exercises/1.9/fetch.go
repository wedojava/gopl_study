// go run fetch.go http://gopl.io
// https://github.com/ray-g/gopl/tree/master/ch01/ex1.07
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(stdout, "resp.Status %s\n", resp.Status)
		_, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: coping %v. status: %s byte(s)\n", err, stdout)
			os.Exit(1)
		}
	}
}
