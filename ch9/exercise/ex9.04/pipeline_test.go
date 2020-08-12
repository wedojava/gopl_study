package pipeline

import "testing"

func benchmarkPipeline(b *testing.B, stages int) {
	in, out := pipeline(stages)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

func BenchmarkPipeline10(b *testing.B)      { benchmarkPipeline(b, 10) }
func BenchmarkPipeline100(b *testing.B)     { benchmarkPipeline(b, 100) }
func BenchmarkPipeline1000(b *testing.B)    { benchmarkPipeline(b, 1000) }
func BenchmarkPipeline10000(b *testing.B)   { benchmarkPipeline(b, 10000) }
func BenchmarkPipeline100000(b *testing.B)  { benchmarkPipeline(b, 100000) }
func BenchmarkPipeline1000000(b *testing.B) { benchmarkPipeline(b, 1000000) }

// go test -benchmem -bench=. -timeout 3600s -v
// goos: darwin
// goarch: amd64
// pkg: gopl.io/ch9/exercise/ex9.04
// BenchmarkPipeline10
// BenchmarkPipeline10-6        	  634978	      1893 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPipeline100
// BenchmarkPipeline100-6       	   60992	     19918 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPipeline1000
// BenchmarkPipeline1000-6      	    6008	    203776 ns/op	      16 B/op	       0 allocs/op
// BenchmarkPipeline10000
// BenchmarkPipeline10000-6     	     415	   3073245 ns/op	    2393 B/op	      24 allocs/op
// BenchmarkPipeline100000
// BenchmarkPipeline100000-6    	      30	  34995420 ns/op	  331497 B/op	    3453 allocs/op
// BenchmarkPipeline1000000
// BenchmarkPipeline1000000-6   	       1	1747014969 ns/op	568844096 B/op	 2803501 allocs/op
// PASS
// ok  	gopl.io/ch9/exercise/ex9.04	9.001s
