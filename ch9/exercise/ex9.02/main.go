package popcount

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc [256]byte
var pcs [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(pc)
}

// PopCount returns the population count (number of set bits) of x.
// This is the fastest one.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bits) of x by looping.
// This is the slowest one
func PopCountLoop(x uint64) int {
	n := 0
	for i := uint(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))]) // convert make this to be the slowest
	}
	return n
}

// PopCountTable returns the population count (number of set bits) of x by table-lookup.
var initTableOnce sync.Once

func initSync() {
	initTableOnce.Do(func() {
		for i := range pc {
			pcs[i] = pcs[i/2] + byte(i&1)
		}
	})
}

// Here is the best practice, but not the fastest one
func PopCountTableSync(x uint64) int {
	initSync()
	return int(pcs[byte(x>>(0*8))] +
		pcs[byte(x>>(1*8))] +
		pcs[byte(x>>(2*8))] +
		pcs[byte(x>>(3*8))] +
		pcs[byte(x>>(4*8))] +
		pcs[byte(x>>(5*8))] +
		pcs[byte(x>>(6*8))] +
		pcs[byte(x>>(7*8))])
}
