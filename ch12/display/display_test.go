package display

import (
	"os"
	"reflect"
)

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

	Display("os.Stderr", os.Stderr)
	// Output:
	// Display os.Stderr(*os.File):

	Display("rV", reflect.ValueOf(os.Stderr))

	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}
