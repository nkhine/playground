package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, validA := corner(i+1, j)
			bx, by, validB := corner(i, j)
			cx, cy, validC := corner(i, j+1)
			dx, dy, validD := corner(i+1, j+1)
			if validA && validB && validC && validD {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Check if z is NaN or Infinity
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0
	}
	return math.Sin(r) / r
}
