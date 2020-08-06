package memo_test

import (
	"testing"

	"gopl.io/ch9/memo1"
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

// go test -v gopl.io/ch9/memo1
// === RUN   Test
// https://golang.org, 1.154334955s, 7359 bytes
// https://godoc.org, 961.53848ms, 7522 bytes
// https://play.golang.org, 1.36106206s, 6314 bytes
// http://gopl.io, 3.658376441s, 4154 bytes
// https://golang.org, 669ns, 7359 bytes
// https://godoc.org, 259ns, 7522 bytes
// https://play.golang.org, 213ns, 6314 bytes
// http://gopl.io, 248ns, 4154 bytes
// --- PASS: Test (7.14s)
// === RUN   TestConcurrent
// https://golang.org, 275.710949ms, 7359 bytes
// https://golang.org, 282.38117ms, 7359 bytes
// https://godoc.org, 283.249182ms, 7522 bytes
// https://godoc.org, 304.106072ms, 7522 bytes
// http://gopl.io, 564.973794ms, 4154 bytes
// https://play.golang.org, 755.575855ms, 6314 bytes
// https://play.golang.org, 783.088955ms, 6314 bytes
// http://gopl.io, 934.74865ms, 4154 bytes
// --- PASS: TestConcurrent (0.94s)
// PASS
// ok  	gopl.io/ch9/memo1	9.138s
