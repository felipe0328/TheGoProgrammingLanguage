// Exercise 3.2: Experiment with vasualizations of other functiosn from the math package.
// Can you produce an egg box, moguls, or a saddle?

package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var graphs = []func(float64, float64) float64{
	sinGraph,
	eggBoxGraph,
	saddleGraph,
}

func main() {
	graphName := "graphs.html"
	graphFile, err := os.Create(graphName)
	if err != nil {
		fmt.Printf("Unable to create %s: %v\n", graphName, err)
	}
	defer graphFile.Close()

	for _, graph := range graphs {
		drawGraph(graph, graphFile)
	}
}

func drawGraph(graph func(float64, float64) float64, w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width= '%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, graph)
			bx, by := corner(i, j, graph)
			cx, cy := corner(i, j+1, graph)
			dx, dy := corner(i+1, j+1, graph)
			fmt.Fprintf(w, "<polygon points='%g, %g %g, %g %g, %g %g, %g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, graph func(float64, float64) float64) (float64, float64) {
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
	return sx, sy
}

func sinGraph(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return (math.Sin(r) / r)
}

func eggBoxGraph(x, y float64) float64 {
	var a, b float64 = 0.1, math.Pi / 2
	return (a * (math.Sin(x/b) + math.Sin(y/b)))
}

func saddleGraph(x, y float64) float64 {
	var a, b float64 = 17, 10
	return (math.Pow((x/a), 2) - math.Pow((y/b), 2))
}
