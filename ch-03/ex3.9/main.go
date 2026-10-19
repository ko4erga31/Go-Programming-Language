/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.9
 *
 * Task:
 *   Write a web server that renders fractals and writes the image data
 *   to the client. Allow the client to specify the x, y, and zoom values
 *   as parameters to the HTTP request.
 *
 * Run:
 *   go run .
 *   then open http://localhost:8000/Mandelbrot?x=-0.75&y=0.1&zoom=50
 *   optional: width, height (16..2048), zoom (1..1e13)
 * =========================================================================
 */

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

type parametrs struct {
	width      int
	height     int
	x, y, zoom float64
}

func main() {
	http.HandleFunc("/Mandelbrot", httpMandelbrot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawPng(w io.Writer, p *parametrs) {
	var xmin, xmax float64 = p.x - 2/p.zoom, p.x + 2/p.zoom
	var ymin, ymax float64 = p.y - 2/p.zoom, p.y + 2/p.zoom
	dx, dy := (float64(xmax-xmin)/float64(p.width))/4, (float64(ymax-ymin)/float64(p.height))/4

	img := image.NewRGBA(image.Rect(0, 0, p.width, p.height))
	for py := 0; py < p.height; py++ {
		y := float64(py)/float64(p.height)*(ymax-ymin) + ymin
		for px := 0; px < p.width; px++ {
			x := float64(px)/float64(p.width)*(xmax-xmin) + (xmin)
			pp, mm := mandelbrot(complex(x+dx, y+dy)), mandelbrot(complex(x-dx, y-dy))
			pm, mp := mandelbrot(complex(x+dx, y-dy)), mandelbrot(complex(x-dx, y+dy))
			avgC := averageColor(pp, mm, mp, pm)
			img.Set(px, py, avgC)
		}
	}
	png.Encode(w, img)
}

func parseURlParametrsINT(w http.ResponseWriter, r *http.Request, s string, def int) (int, bool) {
	if str := r.URL.Query().Get(s); str != "" {
		value, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, "Error of parse "+s, http.StatusBadRequest)
			return 0, false
		}
		return value, true
	}
	return def, true
}

func parseURlParametrsFLOAT(w http.ResponseWriter, r *http.Request, s string, def float64) (float64, bool) {
	if str := r.URL.Query().Get(s); str != "" {
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			http.Error(w, "Error of parse "+s, http.StatusBadRequest)
			return 0, false
		}
		return value, true
	}
	return def, true
}

func isInRangeINT(w http.ResponseWriter, r *http.Request, s string, value int, left, right int) bool {
	if value < left || value > right {
		http.Error(w, "Value "+s+" is not in range["+strconv.Itoa(left)+","+strconv.Itoa(right)+"]", http.StatusBadRequest)
		return false
	}
	return true
}

func isInRangeFLOAT(w http.ResponseWriter, r *http.Request, s string, value, left, right float64) bool {
	if !(value >= left && value <= right) {
		http.Error(w, "Value "+s+" is not in range["+strconv.FormatFloat(left, 'f', -1, 64)+
			","+strconv.FormatFloat(right, 'f', -1, 64)+"]", http.StatusBadRequest)
		return false
	}
	return true
}

func httpMandelbrot(w http.ResponseWriter, r *http.Request) {
	p := &parametrs{width: 1024, height: 1024, x: 0, y: 0, zoom: 1}
	var valueINT int
	var valueFLOAT float64
	var ok bool

	if valueINT, ok = parseURlParametrsINT(w, r, "width", 1024); !ok {
		return
	}
	if ok = isInRangeINT(w, r, "width", valueINT, 16, 2048); !ok {
		return
	}
	p.width = valueINT

	if valueINT, ok = parseURlParametrsINT(w, r, "height", 1024); !ok {
		return
	}
	if ok = isInRangeINT(w, r, "height", valueINT, 16, 2048); !ok {
		return
	}
	p.height = valueINT

	if valueFLOAT, ok = parseURlParametrsFLOAT(w, r, "x", 0); !ok {
		return
	}
	if ok = isInRangeFLOAT(w, r, "x", valueFLOAT, -2.0, 2.0); !ok {
		return
	}
	p.x = valueFLOAT

	if valueFLOAT, ok = parseURlParametrsFLOAT(w, r, "y", 0); !ok {
		return
	}
	if ok = isInRangeFLOAT(w, r, "y", valueFLOAT, -2.0, 2.0); !ok {
		return
	}
	p.y = valueFLOAT

	if valueFLOAT, ok = parseURlParametrsFLOAT(w, r, "zoom", 1); !ok {
		return
	}
	if ok = isInRangeFLOAT(w, r, "zoom", valueFLOAT, 1, 1e13); !ok {
		return
	}
	p.zoom = valueFLOAT

	w.Header().Set("Content-Type", "image/png")
	drawPng(w, p)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - contrast*n, n + 100, n + 200}
		}
	}
	return color.Black
}

func averageColor(colors ...color.Color) color.Color {
	var sum [4]uint32
	lenC := uint32(len(colors))
	if lenC == 0 {
		return color.Black
	}
	for _, c := range colors {
		r, g, b, a := c.RGBA()
		sum[0] += r
		sum[1] += g
		sum[2] += b
		sum[3] += a
	}
	return color.RGBA64{uint16(sum[0] / lenC), uint16(sum[1] / lenC),
		uint16(sum[2] / lenC), uint16(sum[3] / lenC)}
}
