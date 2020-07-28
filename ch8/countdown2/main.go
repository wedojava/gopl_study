package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	subtle()
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // hold until read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	select {
	case t := <-time.After(10 * time.Second): // wait 10 second then send the time to channel
		// Do nothing
		fmt.Println(t)
	case <-abort: // wait until abort recive byte
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}

func subtle() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
