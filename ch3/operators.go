package main

import (
	"fmt"
)

func main() {
	var (
		x uint8 = 1<<1 | 1<<5
		y uint8 = 1<<1 | 1<<2
	)

	fmt.Printf("%08b\n", x) // 00100010
	fmt.Printf("%08b\n", y) // 00000110

	fmt.Printf("%08b\n", x&y) // 00000100
	fmt.Printf("%08b\n", x|y) // 00100110
	fmt.Printf("%08b\n", x^y) // 00100100

	// x &^ y = x & ^y = 00100010 & 11111001
	//   00100010
	// & 11111001
	// -----------
	//   00100000
	fmt.Printf("%08b\n", x&^y) // 00100000

	for i := 0; i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1" "5"
		}

		fmt.Printf("%08b\n", x<<1) // 01000100
		fmt.Printf("%08b\n", x>>1) // 00010001
	}
}
