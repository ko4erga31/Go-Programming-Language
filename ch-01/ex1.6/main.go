/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.6
 *
 * Task:
 *   Modify lissajous to produce images in multiple colors by adding
 *   more values to palette and then displaying them by changing
 *   the third argument of SetColorIndex in some interesting way.
 *
 * Run:
 *   go run main.go > out.gif
 * =========================================================================
 */

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x0B, 0x3D, 0x2E, 0xFF},
	color.RGBA{0x1A, 0x7F, 0x5A, 0xFF},
	color.RGBA{0x33, 0xFF, 0x99, 0xFF},
	color.RGBA{0xB3, 0xFF, 0xD9, 0xFF},
}

const (
	backgroundIndex    = 0 // black background
	darkGreenIndex     = 1 // dark green
	mediumGreenIndex   = 2 // medium green
	phosphorGreenIndex = 3 // phosphor green
	glowIndex          = 4 // bright glow
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	lenPalette := len(palette)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8(i%(lenPalette-1) + 1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
