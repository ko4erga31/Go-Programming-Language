/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.2
 *
 * Task:
 *   Experiment with visual effects by varying the function f in the
 *   surface plot from exercise 3.1. Can you get images like an
 *   egg box, a saddle, or a hill?
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
	drawingSvg()
}

func drawingSvg() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, af := corner(i+1, j)
			bx, by, bf := corner(i, j)
			cx, cy, cf := corner(i, j+1)
			dx, dy, df := corner(i+1, j+1)
			if af && bf && cf && df {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := choiceF(x, y) // logic of func choice
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
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
