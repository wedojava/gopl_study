package anagrams

import "testing"

func TestFunction(t *testing.T) {
	tcs := []struct {
		params []string
		expect bool
	}{
		{[]string{"abc", "cba"}, true},
		{[]string{"abc", "acb"}, true},
		{[]string{"abc", "bca"}, true},
		{[]string{"abc", "bcd"}, false},
	}

	for _, tc := range tcs {
		actual := anagrams(tc.params[0], tc.params[1])
		if actual != tc.expect {
			t.Errorf("Actual: %v, Excepted: %v", actual, tc.expect)
		}
	}
}
