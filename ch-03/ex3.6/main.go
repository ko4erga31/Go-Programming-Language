/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.6
 *
 * Task:
 *   Supersampling is a technique to reduce the effect of pixelation
 *   by computing the color value at several points within each pixel
 *   and taking the average. The simplest method is to divide each pixel
 *   into four subpixels. Implement it.
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
		width, height          = 1024, 1024
		dx, dy                 = (float64(xmax-xmin) / width) / 4, (float64(ymax-ymin) / height) / 4
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			pp, mm := mandelbrot(complex(x+dx, y+dy)), mandelbrot(complex(x-dx, y-dy))
			pm, mp := mandelbrot(complex(x+dx, y-dy)), mandelbrot(complex(x-dx, y+dy))
			avgC := averageColor(pp, mm, mp, pm)
			img.Set(px, py, avgC)
		}
	}
	png.Encode(os.Stdout, img)
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
