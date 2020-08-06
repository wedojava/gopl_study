package memo_test

import (
	"testing"

	"gopl.io/ch9/memo3"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe! Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

// go test -run=TestConcurrent -race -v gopl.io/ch9/memo3
// === RUN   TestConcurrent
// https://golang.org, 872.26383ms, 7359 bytes
// https://golang.org, 875.345242ms, 7359 bytes
// https://godoc.org, 887.132861ms, 7522 bytes
// https://godoc.org, 1.021670614s, 7522 bytes
// https://play.golang.org, 1.341499231s, 6314 bytes
// https://play.golang.org, 1.352604773s, 6314 bytes
// http://gopl.io, 2.295578257s, 4154 bytes
// http://gopl.io, 2.295506493s, 4154 bytes
// --- PASS: TestConcurrent (2.30s)
// PASS
// ok  	gopl.io/ch9/memo3	3.023s
