package sexpr

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		log.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("MarshalIndent() = %s\n", data)
}

func TestFloat(t *testing.T) {
	var tcs = []struct {
		num32 float32
		num64 float64
	}{
		{12.3, 3.21},
		{0.0, 10000},
	}

	for _, tc := range tcs {
		got32, err := Marshal(tc.num32)
		if err != nil {
			log.Fatalf("return err %v", err.Error())
		}
		got64, err := Marshal(tc.num64)
		if err != nil {
			log.Fatalf("return err %v", err.Error())
		}
		if string(got32) != fmt.Sprintf("%4.4f", tc.num32) {
			t.Errorf("Got = %v, Want = %v", got32, tc.num32)
		}
		if string(got64) != fmt.Sprintf("%4.4f", tc.num64) {
			t.Errorf("Got = %v, Want = %v", got64, tc.num64)
		}
	}
}

func TestComplex(t *testing.T) {
	var tcs = []struct {
		num64  complex64
		num128 complex128
	}{
		{12.0 - 3i, 3.2 + 1i},
		{0.0 - 0i, 10000 + 0i},
	}
	for _, tc := range tcs {
		got64, err := Marshal(tc.num64)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		got128, err := Marshal(tc.num128)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(got64) != fmt.Sprintf("#C(%4.4f %4.4f)", real(tc.num64), imag(tc.num64)) {
			t.Errorf("Got: %s, Want: %v", got64, tc.num64)
		}
		if string(got128) != fmt.Sprintf("#C(%4.4f %4.4f)", real(tc.num128), imag(tc.num128)) {
			t.Errorf("Got: %s, Want: %v", got128, tc.num128)
		}
	}
}

func TestBool(t *testing.T) {
	var tcs = []struct {
		b    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, tc := range tcs {
		got, err := Marshal(tc.b)
		if err != nil {
			log.Fatal(err)
		}
		if string(got) != tc.want {
			t.Errorf("Got: %v, Want: %v", got, tc.want)
		}
	}
}

func TestMarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B    bool `sexpr:"bee"`
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		I    Interface
	}
	tcs := []struct {
		r    Record
		want string
	}{
		{
			Record{true, 2.5, 0, 1 + 2i, 2 + 3i, Interface(5)},
			`((bee t) (F32 2.5) (F64 0) (C64 #C(1 2)) (C128 #C(2 3)) (I ("sexpr.Interface" 5)))`,
		},
		{
			Record{false, 0, 1.5, 0, 1i, Interface(0)},
			`((bee nil) (F32 0) (F64 1.5) (C64 #C(0 0)) (C128 #C(0 1)) (I ("sexpr.Interface" 0)))`,
		},
	}
	for _, tc := range tcs {
		data, err := Marshal(tc.r)
		s := string(data)
		if err != nil {
			t.Errorf("Marshal(%s): %s", s, err)
		}
		if s != tc.want {
			t.Errorf("Marshal(%#v)\ngot :%s\nwant:%s", tc.r, s, tc.want)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B   bool
		F32 float32
		F64 float64
		I   Interface `sexpr:"face"`
	}
	Interfaces["sexpr.Interface"] = reflect.TypeOf(int(0))
	tcs := []struct {
		s    string
		want Record
	}{
		{
			`((B t) (F32 2.5) (F64 0) (I ("sexpr.Interface" 5)))`,
			Record{true, 2.5, 0, Interface(5)},
		},
		{
			`((B nil) (F32 0) (F64 1.5) (face ("sexpr.Interface" 0)))`,
			Record{false, 0, 1.5, Interface(0)},
		},
	}
	for _, tc := range tcs {
		var r Record
		err := Unmarshal([]byte(tc.s), &r)
		if err != nil {
			t.Errorf("Unmarshal(%q): %s", tc.s, err)
		}
		if !reflect.DeepEqual(r, tc.want) {
			t.Errorf("Unmarshal(%q)\n got: %#v\nwant: %#v", tc.s, r, tc.want)
		}
	}

}
