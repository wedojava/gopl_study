// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//var out io.Writer = os.Stdout

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	queryName := "cycles"
	cycles := 5

	if len(r.URL.RawQuery) != 0 {
		queries, err := url.ParseQuery(r.URL.RawQuery)
		if err == nil {
			if v, ok := queries[queryName]; ok {
				cycles, _ = strconv.Atoi(v[0])
			}
		}
	}

	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	palette := []color.Color{color.White, color.Black}
	const (
		whiteIndex = 0      // first color in palette
		blackIndex = 1      // next color in palette
		res        = 0.0001 // angular revolution 角转
		size       = 100    // image canvas covers [-size .. +size] 图片画布封面[-size .. + size]
		nframes    = 64     // number of animation frames 动画帧数
		delay      = 8      // delay between frames in 10ms units 帧之间的延迟，以10ms为单位
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference 相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-
