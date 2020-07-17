package targetlist

import (
	"fmt"

	// "gopl.io/ch10/pkglevel1/func1"
	"gopl.io/ch10/pkglevel1/func2"
)

func TargetList() {
	// func1.Func1()
	func2.Func2()
	fmt.Println("hi, I'm `pkglevel1/func1/targetlist/TargetList`")
}
