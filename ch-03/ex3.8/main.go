/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.8
 *
 * Task:
 *   Rendering fractals at high zoom levels demands great arithmetic precision.
 *   Implement the same fractal using four different representations of numbers:
 *   complex64, complex128, big.Float, and big.Rat.
 *   How do they compare in performance and memory usage?
 *   At what zoom levels do rendering artifacts become visible?
 *
 * Run:
 *   go run . -type=complex128 > out.png
 *     types: complex64, complex128, bigfloat, bigrat
 *     (bigrat never finishes with 200 iterations, see NB)
 *   go test -v                    all types give the same colors
 *   go test -bench=. -benchmem    speed and memory, see NB
 *
 * NB: time and memory for one point escaping after 10 iterations
 *     (go test -bench=. -benchmem);
 *
 *   type        time/point   memory/point
 *   complex64      67 ns          3 B
 *   complex128     89 ns          3 B
 *   big.Float     5.2 us        1.2 KB
 *   big.Rat       107 ms        2.7 MB
 * =========================================================================
 */

package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

var (
	funcFlag = flag.String("type", "complex128", "number type: complex64, complex128, bigfloat, bigrat")
	funcs    = map[string]func(float64, float64) color.Color{
		"complex64":  mandelbrotC64,
		"complex128": mandelbrotC128,
		"bigfloat":   mandelbrotBigFloat,
		"bigrat":     mandelbrotBigRat,
	}
)

func main() {
	flag.Parse()
	f, ok := funcs[*funcFlag]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown type %q\n", *funcFlag)
		os.Exit(1)
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, f(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

const (
	iterations             = 200
	contrast               = 15
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 256, 256
)

func mandelbrotC128(x, y float64) color.Color {
	z := complex(x, y)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - contrast*n, n + 100, n + 200}
		}
	}
	return color.Black
}

func mandelbrotC64(x, y float64) color.Color {
	z := complex(float32(x), float32(y))
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if real(v)*real(v)+imag(v)*imag(v) > 4 {
			return color.YCbCr{255 - contrast*n, n + 100, n + 200}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(x, y float64) color.Color {
	zRe := new(big.Float).SetPrec(256).SetFloat64(x)
	zImg := new(big.Float).SetPrec(256).SetFloat64(y)
	vRe, vImg := new(big.Float).SetPrec(256), new(big.Float).SetPrec(256)
	aa, bb := new(big.Float).SetPrec(256), new(big.Float).SetPrec(256)
	sum, ab2 := new(big.Float).SetPrec(256), new(big.Float).SetPrec(256)
	const4 := new(big.Float).SetInt64(4)
	// v*v+z = (a+bi)(a+bi)+(c+di) =(a*a+c-b*b)+(2abi+di)
	for n := uint8(0); n < iterations; n++ {
		ab2.Mul(vRe, vImg).Add(ab2, ab2)
		vRe.Add(aa, zRe).Sub(vRe, bb)
		vImg.Add(ab2, zImg)
		aa.Mul(vRe, vRe)
		bb.Mul(vImg, vImg)
		sum.Add(aa, bb)
		if sum.Cmp(const4) > 0 {
			return color.YCbCr{255 - contrast*n, n + 100, n + 200}
		}
	}
	return color.Black
}

func mandelbrotBigRat(x, y float64) color.Color {
	zRe := new(big.Rat).SetFloat64(x)
	zImg := new(big.Rat).SetFloat64(y)
	vRe, vImg := new(big.Rat), new(big.Rat)
	aa, bb := new(big.Rat), new(big.Rat)
	sum, ab2 := new(big.Rat), new(big.Rat)
	const4 := new(big.Rat).SetInt64(4)
	for n := uint8(0); n < iterations; n++ {
		ab2.Mul(vRe, vImg).Add(ab2, ab2)
		vRe.Add(aa, zRe).Sub(vRe, bb)
		vImg.Add(ab2, zImg)
		aa.Mul(vRe, vRe)
		bb.Mul(vImg, vImg)
		sum.Add(aa, bb)
		if sum.Cmp(const4) > 0 {
			return color.YCbCr{255 - contrast*n, n + 100, n + 200}
		}
	}
	return color.Black
}
