// Package memo provides a concurrency-safe
// memoization of a function of type Func.
// Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import "sync"

// A Memo caches the results of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get is concurrency-safe!
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is first request for this key.
		// This goroutine becomes reponsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})} // init e
		memo.cache[key] = e                    // init cache
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key) // assign value and err to e

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
		// if it is not ready, block here wait other goroutine get and fill &entry,
		// &entry filled means e filled, means memo.cache[key] filled.
		// and e.ready closed. so <-e.ready always recive 0, all e.ready blocks passed.
		// ifelse completed, e.res returned.
	}
	return e.res.value, e.res.err
}

// Our concurrent, duplicate-suppressing, non-blocking cache is complete.
