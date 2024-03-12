// Exercise 3.7: Another simple fractal uses Newton's method to find complex solutions to a function such as
// z^4-1=0. Shade each starting point by the number of iterations required to get close to one of the four roots.
// Color each point by the root it approaches.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const tolerance = 0.0000001

type colorObject struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var (
	roots = []complex128{
		complex(1, 0),
		complex(-1, 0),
		complex(0, 1),
		complex(0, -1),
	}

	rootColors = []colorObject{
		{R: 255, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 255, A: 255},
		{R: 0, G: 255, B: 0, A: 255},
		{R: 255, B: 0, G: 255, A: 255},
	}
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			z := complex(x, y)

			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		z -= zFunction(z) / zDerivative(z)

		for i, root := range roots {
			difference := z - root
			if cmplx.Abs(difference) < tolerance {
				return rootColors[i].GetColor(contrast * n)
			}
		}
	}
	return color.Black
}

// z^4-1
func zFunction(z complex128) complex128 {
	return cmplx.Pow(z, 4) - complex(1, 0)
}

// 4z^3
func zDerivative(z complex128) complex128 {
	return complex(4, 0) * cmplx.Pow(z, 3)
}

func (c *colorObject) GetColor(contrast uint8) color.Color {
	return color.RGBA{
		R: c.GetContrast(c.R, contrast),
		G: c.GetContrast(c.G, contrast),
		B: c.GetContrast(c.B, contrast),
		A: c.A,
	}
}

func (c *colorObject) GetContrast(color, contrast uint8) uint8 {
	if color == 0 {
		return 0
	}

	if color < contrast {
		return contrast
	}

	return color - contrast
}
