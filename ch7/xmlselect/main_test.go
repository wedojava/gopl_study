package main

import (
	"testing"
)

func TestContainsAll(t *testing.T) {
	tcs := []struct {
		src  []string
		sub  []string
		want bool
	}{
		{[]string{"html", "body", "div", "div", "h1"}, []string{"html", "div", "h1"}, true},
	}
	for _, tc := range tcs {
		got := containsAll(tc.src, tc.sub)
		if !got {
			t.Errorf("want: %v, got: %v", tc.want, got)
		}
	}
}
