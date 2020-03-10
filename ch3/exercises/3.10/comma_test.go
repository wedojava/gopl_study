package comma2

import "testing"

func TestComma(t *testing.T) {
	cases := map[string]string{
		"":        "",
		"1":       "1",
		"123":     "123",
		"1234":    "1,234",
		"123456":  "123,456",
		"1234567": "1,234,567",
	}
	for k, v := range cases {
		actual := comma(k)
		if actual != v {
			t.Errorf("Expected: %v, but actual: %v", v, actual)
		}
	}
}
