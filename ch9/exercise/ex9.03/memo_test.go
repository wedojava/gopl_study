package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done <-chan bool) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	key := "https://golang.org"
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		v, err := m.Get(key, nil)
		wg1.Done()
		if v == nil {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil, err)
		}
	}()
	wg1.Wait()

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		done := make(chan bool)
		close(done)
		v, err := m.Get(key, done)
		if v != nil || err == nil {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil, "cancled")
		}
		wg2.Done()
	}()
	wg2.Wait()
}

// go test
// https://golang.org, 1.288683603s, 7357 bytes
// https://godoc.org, 2.384097906s, 7522 bytes
// https://play.golang.org, 1.839392069s, 6314 bytes
// http://gopl.io, 7.469410039s, 4154 bytes
// https://golang.org, 5.869µs, 7357 bytes
// https://godoc.org, 2.691µs, 7522 bytes
// https://play.golang.org, 2.577µs, 6314 bytes
// http://gopl.io, 2.051µs, 4154 bytes
// https://golang.org, 282.159308ms, 7357 bytes
// https://golang.org, 282.165141ms, 7357 bytes
// https://godoc.org, 460.012123ms, 7522 bytes
// https://godoc.org, 459.998389ms, 7522 bytes
// https://play.golang.org, 782.364177ms, 6314 bytes
// https://play.golang.org, 782.291101ms, 6314 bytes
// http://gopl.io, 955.75384ms, 4154 bytes
// http://gopl.io, 955.656183ms, 4154 bytes
// PASS
// ok  	gopl.io/ch9/exercise/ex9.03	16.352s

// go test -run=TestCancel
// PASS
// ok  	gopl.io/ch9/exercise/ex9.03	1.367s
