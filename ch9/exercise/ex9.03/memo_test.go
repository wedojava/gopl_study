package memo_test

import (
	"testing"

	memo "gopl.io/ch9/exercise/ex9.03"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe! Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

// go test -run=TestConcurrent -race -v gopl.io/ch9/memo5
// === RUN   TestConcurrent
// https://godoc.org, 851.863413ms, 7522 bytes
// https://godoc.org, 852.22799ms, 7522 bytes
// https://golang.org, 862.925321ms, 11298 bytes
// https://golang.org, 863.731648ms, 11298 bytes
// https://play.golang.org, 1.331234063s, 6314 bytes
// https://play.golang.org, 1.330243492s, 6314 bytes
// http://gopl.io, 2.457548797s, 4154 bytes
// http://gopl.io, 2.456314205s, 4154 bytes
// --- PASS: TestConcurrent (2.46s)
// PASS
// ok	gopl.io/ch9/memo5   2.482s

// go test
// https://golang.org, 977.655771ms, 11298 bytes
// https://godoc.org, 1.015557021s, 7522 bytes
// https://play.golang.org, 1.432728503s, 6314 bytes
// http://gopl.io, 2.874637569s, 4154 bytes
// https://golang.org, 9.556µs, 11298 bytes
// https://godoc.org, 15.94µs, 7522 bytes
// https://play.golang.org, 11.753µs, 6314 bytes
// http://gopl.io, 10.162µs, 4154 bytes
// https://godoc.org, 269.067261ms, 7522 bytes
// https://godoc.org, 269.29707ms, 7522 bytes
// https://golang.org, 283.143443ms, 11298 bytes
// https://golang.org, 283.251154ms, 11298 bytes
// http://gopl.io, 633.561869ms, 4154 bytes
// http://gopl.io, 633.616549ms, 4154 bytes
// https://play.golang.org, 1.076823934s, 6314 bytes
// https://play.golang.org, 1.076942205s, 6314 bytes
// PASS
// ok	gopl.io/ch9/memo5   7.385s
