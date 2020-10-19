package params

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func isVisaNumber(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("must be a string: %v", v)
	}
	re := regexp.MustCompile("^4[0-9]{12}([0-9]{3})?$")
	if !re.MatchString(s) {
		return fmt.Errorf("not a visa number: %q", s)
	}
	return nil
}

func TestVisa(t *testing.T) {
	tcs := []struct {
		req  *http.Request
		want struct {
			Visa string `http:"visa" validate:"visa"`
		}
	}{
		{
			&http.Request{
				Form: url.Values{
					"visa": []string{"5123456789012123"},
				},
			},
			struct {
				Visa string "http:\"visa\" validate:\"visa\""
			}{
				"5123456789012123",
			},
		},
		{
			&http.Request{
				Form: url.Values{
					"visa": []string{"4123456789012123"},
				},
			},
			struct {
				Visa string "http:\"visa\" validate:\"visa\""
			}{
				"4123456789012123",
			},
		},
	}
	for _, tc := range tcs {
		var got struct {
			Visa string `http:"visa" validate:"visa"`
		}
		err := Unpack(tc.req, &got, map[string]Validator{
			"visa": isVisaNumber,
		})
		if err != nil && strings.Contains(err.Error(), "not a visa number:") {
			t.Logf("error visa number: %s catched, test pass.", tc.req.Form["visa"][0])
		}
		if err == nil && !reflect.DeepEqual(tc.want, got) {
			t.Errorf("Unpack(%v)\ngot :%q\nwant:%q", tc.req, got, tc.want)
		}
	}
}

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
