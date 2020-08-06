package memo_test

import (
	"testing"

	"gopl.io/ch9/memo2"
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

//  go test -run=TestConcurrent -race -v gopl.io/ch9/memo2
// === RUN   TestConcurrent
// https://golang.org, 1.171288549s, 7359 bytes
// https://godoc.org, 2.637534196s, 7522 bytes
// https://play.golang.org, 4.178153957s, 6314 bytes
// http://gopl.io, 7.15494815s, 4154 bytes
// https://golang.org, 7.154891678s, 7359 bytes
// https://godoc.org, 7.154731462s, 7522 bytes
// https://play.golang.org, 7.154658407s, 6314 bytes
// http://gopl.io, 7.154561341s, 4154 bytes
// --- PASS: TestConcurrent (7.16s)
// PASS
// ok  	gopl.io/ch9/memo2	7.899s
