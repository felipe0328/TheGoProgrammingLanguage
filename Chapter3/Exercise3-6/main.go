// Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by computing the color value at several
// points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels"
// Implement it.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type colorObject struct {
	R int
	G int
	B int
	A int
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		offsetX                = (xmax - xmin) / width
		offsetY                = (ymax - ymin) / height
	)

	offX := []float64{-offsetX, offsetX}
	offY := []float64{-offsetY, offsetY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			var zColor colorObject

			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					yi := y + offY[i]
					xi := x + offX[j]
					zi := complex(xi, yi)
					zColor.Sum(mandelbrot(zi))
				}
			}

			zColor.ComputeMean(4)

			// Image point (px, py) represents complex value z
			img.Set(px, py, zColor.GetColor())
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrat = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 255 - constrat*n, G: 255 - constrat*n, B: 255, A: 255}
		}
	}
	return color.RGBA{R: 255, G: 0, B: 125, A: 255}
}

func (c *colorObject) Sum(a color.Color) {
	r1, g1, b1, a1 := a.RGBA()

	c.R += int(r1)
	c.G += int(g1)
	c.B += int(b1)
	c.A += int(a1)
}

func (c *colorObject) ComputeMean(objects int) {
	c.R /= objects
	c.G /= objects
	c.B /= objects
	c.A /= objects
}

func (c *colorObject) GetColor() color.Color {
	return color.RGBA{R: uint8(c.R), G: uint8(c.G), B: uint8(c.B), A: uint8(c.A)}
}
