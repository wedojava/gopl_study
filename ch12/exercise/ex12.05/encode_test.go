package json

import (
	"bytes"
	"encoding/json"
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
	t.Logf("\nMarshal() is below:\n%s\n", data)
	fmt.Println(string(data))

	// Decode it
	var movie Movie
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() is below:\n%+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatalf("not equal")
	}
}

func TestFloat32(t *testing.T) {
	var tcs = []struct {
		v    float32
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}

	for _, tc := range tcs {
		data, err := Marshal(tc.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", tc.v, err)
		}
		if string(data) != tc.want {
			t.Errorf("Marshal(%v) got %s, want %s", tc.v, data, tc.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	var tcs = []struct {
		v    float64
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}

	for _, tc := range tcs {
		data, err := Marshal(tc.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", tc.v, err)
		}
		if string(data) != tc.want {
			t.Errorf("Marshal(%v) got %s, want %s", tc.v, data, tc.want)
		}
	}
}

func TestBool(t *testing.T) {
	var tcs = []struct {
		v    bool
		want string
	}{
		{true, "true"},
		{false, "false"},
	}
	for _, tc := range tcs {
		got, err := Marshal(tc.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", tc.v, err)
		}
		if string(got) != tc.want {
			t.Errorf("Marshal(%v) got: %v, want: %v", tc.v, got, tc.want)
		}
	}
}
