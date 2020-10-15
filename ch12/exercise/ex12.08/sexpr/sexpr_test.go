package sexpr

import (
	"bytes"
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
	t.Logf("\nMarshal() = %s\n", data)

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("\nMarshalIndent() is below:\n%s\n", data)
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

func TestStreamingDecode(t *testing.T) {
	type Book struct {
		Title, Author string
	}
	book := Book{"Point Counterpoint", "Aldous Huxley"}
	data, err := Marshal(book) // encode book to S-expression, assign to data
	if err != nil {
		t.Errorf("setting up test: %s", err)
		return
	}
	data = bytes.Repeat(data, 2) // make test case easy to test More func
	t.Logf("%s", data)

	dec := NewDecoder(bytes.NewReader(data))
	books := make([]Book, 0)
	for dec.More() {
		var b Book
		err := dec.Decode(&b)
		if err != nil {
			t.Errorf("Error after decoding %s: %s", books, err)
			return
		}
		books = append(books, b)
	}
	want := []Book{book, book}
	if !reflect.DeepEqual(want, books) {
		t.Errorf("Got %s, want %s", books, want)
	}
}
