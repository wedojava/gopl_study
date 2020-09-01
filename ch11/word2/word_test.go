package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"的地的", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palidrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// go test -bench=.
// goos: linux
// goarch: amd64
// pkg: gopl.io/ch11/word2
// BenchmarkIsPalindrome-4          4866424               240 ns/op
// PASS
// ok      gopl.io/ch11/word2      1.424s
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// go test -bench=. -benchmem
// goos: linux
// goarch: amd64
// pkg: gopl.io/ch11/word2
// BenchmarkIsPalindrome-4          4865000               240 ns/op             248 B/op          5 allocs/op
// BenchmarkIsPalindrome2-4         9546078               124 ns/op             128 B/op          1 allocs/op
// PASS
// ok      gopl.io/ch11/word2      2.739s
func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome2("A man, a plan, a canal: Panama")
	}
}
