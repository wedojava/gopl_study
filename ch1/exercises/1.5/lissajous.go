package main

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0x00,
		A: 0xff,
	},
	color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5      // number of complete x oscillator revolutions 完整的x振荡器转数
		res     = 0.0001 // angular revolution 角转
		size    = 100    // image canvas covers [-size .. +size] 图片画布封面[-size .. + size]
		nframes = 64     // number of animation frames 动画帧数
		delay   = 8      // delay between frames in 10ms units 帧之间的延迟，以10ms为单位
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference 相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	//gif.EncodeAll(out, &anim)

	buf := new(bytes.Buffer)
	err := gif.EncodeAll(buf, &anim)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("test.gif", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
