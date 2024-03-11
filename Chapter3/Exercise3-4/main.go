// Exercise 3.4: Following the approach of the lissajous example in Section 1.7, construct a web server that computes
// surfaces and writes SVG data to the client.

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	red           = 0xFF0000            // Red Color
	blue          = 0x0000FF            // Blue Color
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", svgHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func svgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	color := r.URL.Query().Get("color")

	maxColor := red

	if color != "" {
		parsedColor, err := strconv.Atoi(color)
		if err != nil {
			fmt.Printf("Error converting color %s: %v", color, err)
		}
		maxColor = parsedColor
	}

	drawGraph(sinGraph, w, maxColor)
}

func drawGraph(graph func(float64, float64) float64, w io.Writer, bigColor int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width= '%d' height='%d'>\n", width, height)

	minZ, maxZ := math.MaxFloat64, -math.MaxFloat64

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az := corner(i+1, j, graph)
			_, _, bz := corner(i, j, graph)
			_, _, cz := corner(i, j+1, graph)
			_, _, dz := corner(i+1, j+1, graph)
			maxVal := max(az, bz, cz, dz)
			minVal := min(az, bz, cz, dz)

			minZ = min(minZ, minVal)
			maxZ = max(maxZ, maxVal)
		}
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, graph)
			bx, by, bz := corner(i, j, graph)
			cx, cy, cz := corner(i, j+1, graph)
			dx, dy, dz := corner(i+1, j+1, graph)

			z := mean(az, bz, cz, dz)
			zPercentage := heightPercentage(z, minZ, maxZ)
			color := getColor(zPercentage, bigColor)

			fmt.Fprintf(w, "<polygon points='%g, %g %g, %g %g, %g %g, %g' style='fill:#%06x'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, graph func(float64, float64) float64) (float64, float64, float64) {
	// find point (x,y) at corner of cells (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := graph(x, y)
	if math.IsNaN(z) {
		z = 0
	}

	if math.IsInf(z, 1) {
		z = math.MaxFloat64
	}
	if math.IsInf(z, -1) {
		z = -math.MaxFloat64
	}

	// Project (x,y,z), isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func sinGraph(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return (math.Sin(r) / r)
}

func heightPercentage(z, min, max float64) float64 {
	currentZ := z - min // we move the data to get the min as zero
	possibleMax := max - min
	return currentZ / possibleMax
}

func getColor(currentHeight float64, bigColor int) int {
	borders := bigColor - blue
	borderData := float64(borders) * currentHeight
	final := blue + borderData
	return int(final)
}

func max(data ...float64) float64 {
	max := -math.MaxFloat64

	for _, val := range data {
		if val > max {
			max = val
		}
	}

	return max
}

func min(data ...float64) float64 {
	min := math.MaxFloat64

	for _, val := range data {
		if val < min {
			min = val
		}
	}

	return min
}

func mean(data ...float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64

	for _, value := range data {
		sum += value
	}

	return sum / float64(len(data))
}
