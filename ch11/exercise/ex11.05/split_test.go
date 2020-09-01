package split

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want []string
	}{
		{s: "", sep: ",", want: []string{""}},
		{s: "premature abstraction:", sep: " ",
			want: []string{"premature", "abstraction:"}},
		{s: "1,2,3,4,5", sep: ",",
			want: []string{"1", "2", "3", "4", "5"}},
		{s: "www.google.com", sep: ".",
			want: []string{"www", "google", "com"}},
		{s: "hello, world", sep: "o",
			want: []string{"hell", ", w", "rld"}},
		{s: "你好，世界", sep: "，",
			want: []string{"你好", "世界"}},
	}

	for _, test := range tests {
		if got := strings.Split(test.s, test.sep); !reflect.DeepEqual(test.want, got) {
			t.Errorf("Split(%q, %q) returned %v, want %v", test.s, test.sep, got, test.want)
		}
	}
}
