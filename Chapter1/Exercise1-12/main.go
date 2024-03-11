// Modify the Lissajous server to read parameter values from the URL.
// For example, you might arrange it so that a URL like http:localhost:8000/?cycles=20
// sets the number of cycles to 20 instead of the default 5.
// Use strconv.Atoi function to convert the string parameter into a integer.
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
	"strconv"
)

var palette = []color.Color{color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.Black}

const (
	greenIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	headerQuery := r.URL.Query().Get("cycles")
	cycles, _ := strconv.Atoi(headerQuery)

	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	if cycles == 0 {
		cycles = 5
	}

	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas cover [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
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