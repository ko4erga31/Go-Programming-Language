/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.3
 *
 * Task:
 *   Color each polygon with a color that depends on its height,
 *   so that the peaks are red (#ff0000) and the bottoms are blue (#0000ff).
 *
 * Run:
 *   go run . -f saddle > out.svg
 * =========================================================================
 */

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var (
	funcFlag = flag.String("f", "sin(x)/x", "type of function:\n\teggBox\n\tsaddle\n\thill\n")
	funcs    = map[string]func(float64, float64) float64{
		"sin(x)/x": f,
		"eggBox":   eggBox,
		"saddle":   saddle,
		"hill":     hill,
	}
	choiceF func(float64, float64) float64
)

func main() {
	flag.Parse()
	var ok bool
	choiceF, ok = funcs[*funcFlag]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown function: %q\n", *funcFlag)
		os.Exit(1)
	}
	drawingSVG()
}

func minMaxSVG() (float64, float64) {
	var zmin, zmax float64 = math.Inf(1), math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az, af := corner(i+1, j)
			_, _, bz, bf := corner(i, j)
			_, _, cz, cf := corner(i, j+1)
			_, _, dz, df := corner(i+1, j+1)

			if af && bf && cf && df {
				zmin = min(zmin, az, bz, cz, dz)
				zmax = max(zmax, az, bz, cz, dz)
			}
		}
	}
	return zmin, zmax
}

func fillFugureSVG(zmin, zmax float64) {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, af := corner(i+1, j)
			bx, by, bz, bf := corner(i, j)
			cx, cy, cz, cf := corner(i, j+1)
			dx, dy, dz, df := corner(i+1, j+1)

			if af && bf && cf && df {
				zavg := (az + bz + cz + dz) / 4
				color := colorFor(zmin, zmax, zavg)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
}

func drawingSVG() {

	printHeaderSVG()

	zmin, zmax := minMaxSVG()

	fillFugureSVG(zmin, zmax)

	fmt.Println("</svg>")
}

func printHeaderSVG() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
}

func colorFor(zmin, zmax, zavg float64) string {
	t := 0.5
	if zmax != zmin {
		t = (zavg - zmin) / (zmax - zmin)
	}
	red := uint8(t * 255)
	blue := 255 - red
	return fmt.Sprintf("#%02x00%02x", red, blue)
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := choiceF(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 4
}

func saddle(x, y float64) float64 {
	return (x*x - y*y) / 500
}

func hill(x, y float64) float64 {
	return math.Exp(-(x*x + y*y) / 80)
}
