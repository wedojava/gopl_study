// go run main.go -temp 135℉
// 57.22222222222222℃
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
