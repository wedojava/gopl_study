// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
// goroutine 1 [running]:
// main.printStack()
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:58 +0x5b
// panic(0x4a8fe0, 0x55f890)
//         /usr/lib/go/src/runtime/panic.go:969 +0x166
// main.f(0x0)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:51 +0x1dc
// main.f(0x1)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:53 +0x17d
// main.f(0x2)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:53 +0x17d
// main.f(0x3)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:53 +0x17d
// main.main()
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:47 +0x4c
// panic: runtime error: integer divide by zero
//
// goroutine 1 [running]:
// main.f(0x0)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:51 +0x1dc
// main.f(0x1)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:53 +0x17d
// main.f(0x2)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:53 +0x17d
// main.f(0x3)
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:53 +0x17d
// main.main()
//         /home/faceless/go/src/gopl.io/ch5/defer2/main.go:47 +0x4c
// exit status 2

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
