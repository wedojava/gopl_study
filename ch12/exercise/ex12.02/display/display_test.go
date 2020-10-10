package display

import (
	"io"
	"net"
	"os"
	"reflect"
	"testing"

	"gopl.io/ch7/eval"
)

func TestMapKeys(t *testing.T) {
	s := map[struct{ x int }]int{
		{1}: 11,
		{2}: 22,
	}
	Display("s", s)
	// Output:

	// new func added before:
	// Display s (map[struct { x int }]int):
	// s[struct { x int }value] = 11
	// s[struct { x int }value] = 22

	// formatKeys added:
	// s[{x: 2}] = 22
	// s[{x: 1}] = 11

	// map key cannot be the type of slice
	a := map[[3]int]int{
		{1, 2, 3}: 123,
		{2, 3, 4}: 234,
	}
	Display("a", a)
	// Output:

	// formatKeys add before
	// a[[3]intvalue] = 234
	// a[[3]intvalue] = 123

	// formatKeys added
	// Display a (map[[3]int]int):
	// a[2, 3, 4] = 234
	// a[1, 2, 3] = 123
}

func Example_expr() {
	e, _ := eval.Parse("sqrt(A/pi)")
	Display("e", e)
	// Output:
	// Display e (eval.call):
	// e.fn = "sqrt"
	// e.args[0].type = eval.binary
	// e.args[0].value.op = 47
	// e.args[0].value.x.type = eval.Var
	// e.args[0].value.x.value = "A"
	// e.args[0].value.y.type = eval.Var
	// e.args[0].value.y.value = "pi"
}

func Example_slice() {
	Display("slice", []*int{new(int), nil})
	// Output:
	// Display slice ([]*int):
	// (*slice[0]) = 0
	// slice[1] = nil
}

func Example_nilInterface() {
	var w io.Writer
	Display("w", w)
	// Output:
	// Display w (<nil>):
	// w = invalid
}

func Example_ptrToInterface() {
	var w io.Writer
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w) = nil
}

func Example_struct() {
	Display("x", struct{ x interface{} }{3})
	// Output:
	// Display x (struct { x interface {} }):
	// x.x.type = int
	// x.x.value = 3
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}

func Example_ptrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

func Example_array() {
	Display("x", [1]interface{}{3})
	// Output:
	// Display x ([1]interface {}):
	// x[0].type = int
	// x[0].value = 3
}

func Example_movie() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Linoel Mandrake": "Peter Sellers",
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
	Display("strangelove", strangelove)
	// Output:
	// Display strangelove (display.Movie):
	// strangelove.Title = "Dr. Strangelove"
	// strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
	// strangelove.Year = 1964
	// strangelove.Color = false
	// strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
	// strangelove.Actor["Grp. Capt. Linoel Mandrake"] = "Peter Sellers"
	// strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
	// strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
	// strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
	// strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
	// strangelove.Oscars[0] = "Best Actor (Nomin.)"
	// strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
	// strangelove.Oscars[2] = "Best Director (Nomin.)"
	// strangelove.Oscars[3] = "Best Picture (Nomin.)"
	// strangelove.Sequel = nil
}

// Just run `go test` to see the result
func TestWithoutCrashing(t *testing.T) {
	Display("os.Stderr", os.Stderr)

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)

	Display("rV", reflect.ValueOf(os.Stderr))

	// a pointer that points to itself
	type P *P
	var p P
	p = &p
	if false {
		Display("p", p)
		// Output:
		// Display p (display.P):
		// ...stuck, no output...
	}

	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	if true {
		Display("m", m)
		// Output:
		// Display m (display.M):
		// ...stuck, no output...
	}

	// a slice that contains itself
	type S []S
	s := make(S, 1)
	s[0] = s
	if true {
		Display("s", s)
		// Output:
		// Display s (display.S):
		// ...stuck, no output...
	}

	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if true {
		Display("c", c)
		// Output:
		// Display c (display.Cycle):
		// c.Value = 42
		// (*c.Tail).Value = 42
		// (*(*c.Tail).Tail).Value = 42
		// ...ad infinitum...
	}
}
