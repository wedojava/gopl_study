package basename

import "testing"

var tcs = []struct {
	params string
	want   string
}{
	{"a", "a"},
	{"a.go", "a"},
	{"a/b/c.go", "c"},
	{"a/b.c.go", "b.c"},
}

func TestBasename1(t *testing.T) {
	for _, tc := range tcs {
		got := Basename1(tc.params)
		if got != tc.want {
			t.Errorf("want %v, got %v", tc.want, got)
		}
	}
}

func TestBasename2(t *testing.T) {
	for _, tc := range tcs {
		got := basename2(tc.params)
		if got != tc.want {
			t.Errorf("want %v, got %v", tc.want, got)
		}
	}
}
