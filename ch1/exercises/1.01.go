package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	//a := fmt.Sprintf(strings.Join(os.Args[0:], "/"))
	//fmt.Printf("%T", a)
}
