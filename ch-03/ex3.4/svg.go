package main

import (
	"fmt"
	"io"
	"math"
)

type parametrs struct {
	width   int
	height  int
	xyscale float64
	zscale  float64
	choiceF func(float64, float64) float64
	top     color
	bottom  color
}

type color struct {
	red   uint8
	green uint8
	blue  uint8
}

const (
	xyrange = 30.0
	cells   = 100
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var (
	funcs = map[string]func(float64, float64) float64{
		"sin(x)/x": f,
		"eggBox":   eggBox,
		"saddle":   saddle,
		"hill":     hill,
	}
)

func minMaxSVG(p *parametrs) (float64, float64) {
	zmin, zmax := math.Inf(1), math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az, af := corner(i+1, j, p)
			_, _, bz, bf := corner(i, j, p)
			_, _, cz, cf := corner(i, j+1, p)
			_, _, dz, df := corner(i+1, j+1, p)

			if af && bf && cf && df {
				zmin = min(zmin, az, bz, cz, dz)
				zmax = max(zmax, az, bz, cz, dz)
			}
		}
	}
	return zmin, zmax
}

func fillFugureSVG(w io.Writer, zmin, zmax float64, p *parametrs) {
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, af := corner(i+1, j, p)
			bx, by, bz, bf := corner(i, j, p)
			cx, cy, cz, cf := corner(i, j+1, p)
			dx, dy, dz, df := corner(i+1, j+1, p)

			if af && bf && cf && df {
				zavg := (az + bz + cz + dz) / 4
				color := colorFor(zmin, zmax, zavg, p)
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
}

func drawingSVG(w io.Writer, p *parametrs) {
	printHeaderSVG(w, p)
	zmin, zmax := minMaxSVG(p)
	fillFugureSVG(w, zmin, zmax, p)
	fmt.Fprintf(w, "</svg>")
}

func printHeaderSVG(w io.Writer, p *parametrs) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.width, p.height)
}

func colorFor(zmin, zmax, zavg float64, p *parametrs) string {
	t := 0.5
	if zmax != zmin {
		t = (zavg - zmin) / (zmax - zmin)
	}
	red := lerp(p.bottom.red, p.top.red, t)
	green := lerp(p.bottom.green, p.top.green, t)
	blue := lerp(p.bottom.blue, p.top.blue, t)
	return fmt.Sprintf("#%02x%02x%02x", red, green, blue)
}

func lerp(a, b uint8, t float64) uint8 {
	return uint8(float64(a) + t*(float64(b)-float64(a)))
}

func corner(i, j int, p *parametrs) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := p.choiceF(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}
	sx := float64(p.width)/2 + (x-y)*cos30*p.xyscale
	sy := float64(p.height)/2 + (x+y)*sin30*p.xyscale - z*p.zscale
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
