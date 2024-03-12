// Exercise 3.9: Write a web server that render fractals and writes the image data to the client.
// Allow the client to specify the x,y and zoom values as parameters to the http Request
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	offsetX       = 0
	offsetY       = 0
	zoom          = 1 / float64(0.5)
	width, height = 1024, 1024
	iterations    = 200
)

func main() {
	http.HandleFunc("/", fractalHandler)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func fractalHandler(w http.ResponseWriter, r *http.Request) {
	zmS := r.URL.Query().Get("zoom")
	oYS := r.URL.Query().Get("offsetY")
	oXS := r.URL.Query().Get("offsetX")

	zm := zoom
	oy := float64(offsetY)
	ox := float64(offsetX)

	if zmS != "" {
		zmV, err := strconv.ParseFloat(zmS, 64)
		if err != nil {
			fmt.Printf("Unable to parse zoom: %s: %v\n", zmS, err)
		} else {
			zm = 1 / zmV
		}
	}

	if oYS != "" {
		oYV, err := strconv.ParseFloat(oYS, 64)
		if err != nil {
			fmt.Printf("Unable to parse offsetY: %s: %v\n", oYS, err)
		} else {
			oy = oYV
		}
	}

	if oXS != "" {
		oXV, err := strconv.ParseFloat(oXS, 64)
		if err != nil {
			fmt.Printf("Unable to parse offsetX: %s: %v\n", oXS, err)
		} else {
			ox = oXV
		}
	}

	createFractal(w, mandelbrot128, zm, ox, oy)
}

func createFractal(w io.Writer, mandelbrot func(complex128) color.Color, zoom, offsetX, offsetY float64) {
	var xmin, ymin, xmax, ymax = -zoom, -zoom, +zoom, +zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin + offsetY
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin + offsetX
			z := complex(x, y)

			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot128(z complex128) color.Color {
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 255 - contrast*n, G: 255 - contrast*n, B: 255, A: 255}
		}
	}
	return color.RGBA{R: 255, G: 0, B: 0, A: 255}
}
