package main

import (
	"fmt"
	"math"
	"net/http"
)

const (
	width, height = 2000, 900
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Draw() string {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, acz := corner(i+1, j)
			bx, by, bcz := corner(i, j)
			cx, cy, ccz := corner(i, j+1)
			dx, dy, dcz := corner(i+1, j+1)
			svg += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#%x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, uint64((acz+bcz+ccz+dcz)/4))
		}
	}
	svg += "</svg>"
	return svg
}

func corner(i, j int) (float64, float64, uint64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	cz := uint64((0xff00-0xff)*(z+1)/2 + 0xff)
	return sx, sy, ((cz & uint64(0xff00)) << 8) + cz&uint64(0xff)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func server() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(Draw()))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func main() {
	server()
}
