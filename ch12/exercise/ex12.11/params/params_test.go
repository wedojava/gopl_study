package params

import (
	"log"
	"testing"
)

func TestPack(t *testing.T) {
	tcs := []struct {
		s struct {
			Name string `http:"n"`
			Age  int    `http:"a"`
		}
		want string
	}{
		{
			struct {
				Name string "http:\"n\""
				Age  int    "http:\"a\""
			}{"Arugula", 35},
			"a=35&n=Arugula",
		},
	}
	for _, tc := range tcs {
		u, err := Pack(&tc.s)
		if err != nil {
			log.Fatal(err)
		}
		got := u.RawQuery
		if got != tc.want {
			t.Errorf("Pack(%#v): got %q, want %q", tc.s, got, tc.want)
		}
	}
}
