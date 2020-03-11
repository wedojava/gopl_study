package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	// 10000 | 00001 = 10001
	var v Flags = FlagMulticast | FlagUp // "10001"

	// v&FlagUp => 10001 & 00001 = 00001 == FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"

	// *v &^= FlagUp
	// *v = *v & (*v ^ FlagUp)
	// *v = 10001 & (10001 ^ 00001)
	// *v = 10001 & 10000
	// *v = 10000
	TurnDown(&v)

	// v = 10000 (10000 & 00001 = 00000) != FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"

	// *v = *v | FlagBroadcast
	// *v = 10000 | 00010 = 10010
	SetBroadcast(&v)

	// 10010 & 00001 = 0
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10010 false"

	// 10010 & (00010 | 10000) = 10010 & 10010 != 0
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
