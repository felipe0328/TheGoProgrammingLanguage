// Exercise 3.8: Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal
// using four different representations of numbers: complex64, complex128, big.Float and big.Rat. (The latter tow types
// are found in the math/big package. Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-
// precision rational numbers). How do the compare in performance and memory usage?, At what zoom levels do rendering
// artifacts become visilbe?
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	offsetX                = -0.043
	offsetY                = -0.985
	zoom                   = 1 / float64(150)
	xmin, ymin, xmax, ymax = -zoom, -zoom, +zoom, +zoom
	width, height          = 1024, 1024
	iterations             = 200
)

func main() {
	i128, _ := os.Create("f128.png")
	createFractal(i128, mandelbrot128)
	i128.Close()

	i64, _ := os.Create("f64.png")
	createFractal(i64, mandelbrot64)
	i64.Close()

	bigI64, _ := os.Create("bigf64.png")
	createFractal(bigI64, mandelbrotBigFloat)
	bigI64.Close()

	// Really slow
	// bigRat, _ := os.Create("bigRat.png")
	// createFractal(bigRat, mandelbrotBigRat)
	// bigRat.Close()

}

func createFractal(w io.Writer, mandelbrot func(complex128) color.Color) {

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

func mandelbrot64(z complex128) color.Color {
	const contrast = 15

	var z64 = complex64(z)
	var v complex64

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z64
		if cmplx.Abs(complex128(v)) > 2 {
			return color.RGBA{R: 255 - contrast*n, G: 255 - contrast*n, B: 255, A: 255}
		}
	}
	return color.RGBA{R: 255, G: 0, B: 0, A: 255}
}

func mandelbrotBigFloat(z complex128) color.Color {
	const contrast = 15

	r := big.NewFloat(real(z))
	i := big.NewFloat(imag(z))
	var vr, vi big.Float

	comparer := big.NewFloat(4) // 2^2

	for n := 0; n < iterations; n++ {
		// z^2 = (r + ij)^2 = (r+ij)(r+ij) =  (r*r) + 2rij + (ij)^2 = (r*r - i*i) + (2ri)j
		// z^2 + c = ((r^2 - i^2) + (2ri)j) + (cr + cij) = ((r^2 - i^2) + cr) + (2ri + ci)j
		var vr2, vi2 big.Float

		vr2.Mul(&vr, &vr).Sub(&vr2, (&big.Float{}).Mul(&vi, &vi)).Add(&vr2, r)
		vi2.Mul(&vi, &vr).Mul(&vi2, big.NewFloat(2)).Add(&vi2, i)

		vr, vi = vr2, vi2

		if getAbs(vr, vi).Cmp(comparer) == 1 {
			return color.RGBA{R: 255 - contrast*uint8(n), G: 255 - contrast*uint8(n), B: 255, A: 255}
		}
	}
	return color.RGBA{R: 255, G: 0, B: 0, A: 255}
}

func mandelbrotBigRat(z complex128) color.Color {
	const contrast = 15

	r := (&big.Rat{}).SetFloat64(real(z))
	i := (&big.Rat{}).SetFloat64(imag(z))
	var vr, vi big.Rat

	comparer := (&big.Rat{}).SetFloat64(4) // 2^2

	for n := 0; n < iterations; n++ {
		// z^2 = (r + ij)^2 = (r+ij)(r+ij) =  (r*r) + 2rij + (ij)^2 = (r*r - i*i) + (2ri)j
		// z^2 + c = ((r^2 - i^2) + (2ri)j) + (cr + cij) = ((r^2 - i^2) + cr) + (2ri + ci)j
		var vr2, vi2 big.Rat

		vr2.Mul(&vr, &vr).Sub(&vr2, (&big.Rat{}).Mul(&vi, &vi)).Add(&vr2, r)
		vi2.Mul(&vi, &vr).Mul(&vi2, big.NewRat(2, 1)).Add(&vi2, i)

		vr, vi = vr2, vi2

		if getAbsRat(vr, vi).Cmp(comparer) == 1 {
			return color.RGBA{R: 255 - contrast*uint8(n), G: 255 - contrast*uint8(n), B: 255, A: 255}
		}
	}
	return color.RGBA{R: 255, G: 0, B: 0, A: 255}
}

// sqrt(r^2+i^2) => not doing sqrt to omit one opeartion but elevating bot sides by 2
func getAbs(r, i big.Float) *big.Float {
	var result big.Float
	result.Mul(&r, &r)
	result.Add(&result, (&big.Float{}).Mul(&i, &i))

	return &result
}

// sqrt(r^2+i^2) => not doing sqrt to omit one opeartion but elevating bot sides by 2
func getAbsRat(r, i big.Rat) *big.Rat {
	var result big.Rat
	result.Mul(&r, &r)
	result.Add(&result, (&big.Rat{}).Mul(&i, &i))

	return &result
}
