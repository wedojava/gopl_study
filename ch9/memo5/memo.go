// Package memo provides a concurrency-safe
// memoization of a function of type Func.
// Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// a request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

// Memo type consists of a channel, requests, through which the caller of Get
// communicates  with the monitor goroutine.
type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get is concurrency-safe!
// Create a response channel, puts it in the request, sends it to the monitor goroutine
// then immediately receives from it
func (memo *Memo) Get(key string) (value interface{}, err error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	// cache is confined to this monitor goroutine: (*Memo).server
	cache := make(map[string]*entry)
	// Read requests until the request channel is closed by the Close method.
	for req := range memo.requests {
		e := cache[req.key] // consult the cache
		if e == nil {       // create and insert new entry
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
