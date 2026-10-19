/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.7
 *
 * Task:
 *   Another simple fractal uses Newton’s method to find
 *   complex solutions to a function such as z⁴ − 1 = 0.
 *   Shade each starting point by the number of iterations
 *   required to get close to one of the four roots.
 *   Color each point by the root it approaches.
 *
 * Run:
 *   go run . > out.png
 * =========================================================================
 */

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
		dx, dy                 = (float64(xmax-xmin) / width) / 4, (float64(ymax-ymin) / height) / 4
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			pp, mm := newtonSol(complex(x+dx, y+dy)), newtonSol(complex(x-dx, y-dy))
			pm, mp := newtonSol(complex(x+dx, y-dy)), newtonSol(complex(x-dx, y+dy))
			avgC := averageColor(pp, mm, mp, pm)
			img.Set(px, py, avgC)
		}
	}
	png.Encode(os.Stdout, img)
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

func newtonSol(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 8
		epsilon    = 1e-6
		minShade   = 40
	)
	var (
		red    = color.RGBA{0xFF, 0, 0, 0xFF}
		green  = color.RGBA{0, 0xFF, 0, 0xFF}
		blue   = color.RGBA{0, 0, 0xFF, 0xFF}
		yellow = color.RGBA{0xFF, 0xFF, 0, 0xFF}
	)
	solutions := [4]complex128{1, -1, 1i, -1i}
	colors := [4]color.RGBA{red, green, blue, yellow}
	v := z
	for n := uint8(0); n < iterations; n++ {
		v = v - ((v*v*v*v)-1)/(4*(v*v*v))
		for i := range len(solutions) {
			if cmplx.Abs(v-solutions[i]) < epsilon {
				var shade uint32
				approxShade := uint32(contrast) * uint32(n)
				if approxShade > 255-minShade {
					shade = minShade
				} else {
					shade = 255 - approxShade
				}
				r, g, b := colors[i].R, colors[i].G, colors[i].B
				return color.RGBA{uint8(uint32(r) * shade / 255), uint8(uint32(g) * shade / 255), uint8(uint32(b) * shade / 255), 255}
			}
		}
	}
	return color.Black
}
