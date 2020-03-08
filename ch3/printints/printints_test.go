package printints

import "testing"

var tcs = []struct {
	params []int
	want   string
}{
	{[]int{1, 2, 3}, "[1, 2, 3]"},
}

func TestIntsToString(t *testing.T) {
	for _, tc := range tcs {
		got := intsToString(tc.params)
		if got != tc.want {
			t.Errorf("Got %v, want %v", got, tc.want)
		}
	}

}
