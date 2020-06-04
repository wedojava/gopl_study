// now at double func line: 1
// now at the line before return
// now at double func's anonymous func!
// double(4) = 8
// --------------------------
// now at triple begin.
// now in triple below defer but before return
// now at double func line: 1
// now at the line before return
// now at double func's anonymous func!
// double(4) = 8
// triple of 4 is  12

package main

import "fmt"

func main() {
	double(4)
	a := triple(4)
	fmt.Println("triple of 4 is ", a)
}

func double(x int) (result int) {
	fmt.Println("now at double func line: 1")
	defer func() {
		fmt.Println("now at double func's anonymous func!")
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	fmt.Println("now at the line before return")
	return x + x
}

func triple(x int) (result int) {
	fmt.Println("now at triple begin.")
	defer func() { result += x }()
	fmt.Println("now in triple below defer but before return")
	return double(x)
}
