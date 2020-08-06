package memo_test

import (
	"testing"

	"gopl.io/ch9/memo4"
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

// go test -run=TestConcurrent -race -v gopl.io/ch9/memo4
// === RUN   TestConcurrent
// https://godoc.org, 1.3041341s, 7522 bytes
// https://godoc.org, 1.303603352s, 7522 bytes
// https://golang.org, 1.362276673s, 7359 bytes
// https://golang.org, 1.361567814s, 7359 bytes
// https://play.golang.org, 1.363052441s, 6314 bytes
// https://play.golang.org, 1.362413505s, 6314 bytes
// http://gopl.io, 2.300512579s, 4154 bytes
// http://gopl.io, 2.300023193s, 4154 bytes
// --- PASS: TestConcurrent (2.30s)
// PASS
// ok  	gopl.io/ch9/memo4	3.135s
