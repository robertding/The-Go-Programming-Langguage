package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
)

func Draw(f io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(f, img)
}

func supersampling(img *image.NewRGBA) {
	for py :=0; py < img.Rect.Dy(); py++ {
		for px := 0; px < img.Rect.Dx(); px++ {
			color :
		}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 30
	const contrast = 15
	var v complex128
	colors := palette.WebSafe
	len_colors := len(colors)
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colors[int(n)*len_colors/iterations]
		}
	}
	return color.Black
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Draw(w)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}
