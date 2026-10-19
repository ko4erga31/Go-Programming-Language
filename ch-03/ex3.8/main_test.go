package main

import (
	"image/color"
	"testing"
)

var sink color.Color

var points = []struct {
	name string
	x, y float64
}{
	{"n=5", -0.7, 0.6},
	{"n=7", -0.7, 0.5},
	{"n=10", -0.7, 0.4},
}

var types = []string{"complex64", "complex128", "bigfloat", "bigrat"}

func TestSameColor(t *testing.T) {
	for _, p := range points {
		want := mandelbrotC128(p.x, p.y)
		for _, name := range types {
			if got := funcs[name](p.x, p.y); got != want {
				t.Errorf("%s(%v, %v) = %v, want %v", name, p.x, p.y, got, want)
			}
		}
	}
}

// Run: go test -bench=. -benchmem
func BenchmarkMandelbrot(b *testing.B) {
	for _, p := range points {
		for _, name := range types {
			f := funcs[name]
			b.Run(p.name+"/"+name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sink = f(p.x, p.y)
				}
			})
		}
	}
}
