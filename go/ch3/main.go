package main

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
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
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Extract values from query parameters
	q := r.URL.Query()
	customWidth := width
	customHeight := height
	customColor := "grey" // Default color

	// Check for width and height parameters
	if wVal, err := strconv.Atoi(q.Get("width")); err == nil && wVal > 0 {
		customWidth = wVal
	}
	if hVal, err := strconv.Atoi(q.Get("height")); err == nil && hVal > 0 {
		customHeight = hVal
	}

	// Recompute scales
	xyscale := float64(customWidth) / 2 / xyrange
	zscale := float64(customHeight) * 0.4

	// Check for color parameter (and validate it's a color value to avoid XSS issues)
	if colorVal := q.Get("color"); colorVal != "" && isValidColor(colorVal) {
		customColor = colorVal
	}

	fmt.Fprint(w, "<!DOCTYPE html><html><body>")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", customColor, customWidth, customHeight)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1, validA := corner(i+1, j, xyscale, zscale, customWidth, customHeight)
			bx, by, z2, validB := corner(i, j, xyscale, zscale, customWidth, customHeight)
			cx, cy, z3, validC := corner(i, j+1, xyscale, zscale, customWidth, customHeight)
			dx, dy, z4, validD := corner(i+1, j+1, xyscale, zscale, customWidth, customHeight)

			if validA && validB && validC && validD {
				color := computeColor([]float64{z1, z2, z3, z4})
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}

	fmt.Fprint(w, "</svg>")
	fmt.Fprint(w, "</body></html>")
}

func computeColor(heights []float64) string {
	avgHeight := (heights[0] + heights[1] + heights[2] + heights[3]) / 4

	// Map the average height to a color.
	// We'll use a simple linear interpolation between blue and red based on the height.
	clamp := func(v, min, max float64) float64 {
		if v < min {
			return min
		} else if v > max {
			return max
		}
		return v
	}

	scale := clamp((avgHeight+0.5)*0.5, 0, 1)
	r := int(255 * scale)
	b := int(255 * (1 - scale))

	return fmt.Sprintf("#%02x%02x%02x", r, 0, b)
}

// Validates if a string is a valid HTML color (simple check for this example)
func isValidColor(s string) bool {
	// A simple regex to match common color values (hex, named colors, etc.).
	// Note: This is a basic validation and might not capture all valid HTML color values.
	matched, _ := regexp.MatchString(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$|^[a-zA-Z]+$`, s)
	return matched
}
func corner(i, j int, xyscale, zscale float64, customWidth, customHeight int) (float64, float64, float64, bool) {

	// Essentially the same as your `corner` function, but use `xyscale` and `zscale` arguments instead of the constants.
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D
	sx := float64(customWidth)/2 + (x-y)*cos30*xyscale
	sy := float64(customHeight)/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z, true
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0
	}
	return math.Sin(r) / r
}
