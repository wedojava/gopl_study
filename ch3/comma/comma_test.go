package comma

import "testing"

var tcs = []struct {
	params string
	want   string
}{
	{"123456", "123,456"},
	{"123", "123"},
}

func TestComma(t *testing.T) {
	for _, tc := range tcs {
		got := comma(tc.params)
		if got != tc.want {
			t.Errorf("want %v, got %v", tc.want, got)
		}
	}

}
