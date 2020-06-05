// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
// panic: runtime error: integer divide by zero
//
// goroutine 1 [running]:
// main.f(0x0)
//         /home/faceless/go/src/gopl.io/ch5/defer1/main.go:10 +0x1dc
// main.f(0x1)
//         /home/faceless/go/src/gopl.io/ch5/defer1/main.go:12 +0x17d
// main.f(0x2)
//         /home/faceless/go/src/gopl.io/ch5/defer1/main.go:12 +0x17d
// main.f(0x3)
//         /home/faceless/go/src/gopl.io/ch5/defer1/main.go:12 +0x17d
// main.main()
//         /home/faceless/go/src/gopl.io/ch5/defer1/main.go:6 +0x2a
// exit status 2

package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
