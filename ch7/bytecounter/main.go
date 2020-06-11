package main

import "fmt"

type ByteCounter int

// Since c.Write be invoked, c will be set as len(p) after it be converted into ByteCounter
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5 len("hello")

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12 len("hello, Dolly")
}
