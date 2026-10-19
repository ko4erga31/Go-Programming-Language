/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.4
 *
 * Task:
 *   Following the approach of the Lissajous example, construct a web
 *   server that computes surfaces and writes SVG data to the client.
 *   The server must set the Content-Type header to image/svg+xml.
 *   Allow the client to specify values like height, width, and color
 *   as HTTP request parameters.
 *
 * Run:
 *   go run .
 *   then open http://localhost:8000/?f=saddle&width=800&height=400
 * =========================================================================
 */

package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", drawOnSite)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawOnSite(w http.ResponseWriter, r *http.Request) {
	p := &parametrs{width: 600, height: 320, choiceF: funcs["sin(x)/x"], top: color{255, 0, 0}, bottom: color{0, 0, 255}}

	if s := r.URL.Query().Get("f"); s != "" {
		value, exists := funcs[s]
		if !exists {
			http.Error(w, "Error of parse function", http.StatusBadRequest)
			return
		}
		p.choiceF = value
	}
	if s := r.URL.Query().Get("width"); s != "" {
		value, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "Error of parse width", http.StatusBadRequest)
			return
		}
		if value >= 100 && value <= 2000 {
			p.width = value
		} else {
			http.Error(w, "width must be between 100 and 2000", http.StatusBadRequest)
			return
		}
	}
	if s := r.URL.Query().Get("height"); s != "" {
		value, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "Error of parse height", http.StatusBadRequest)
			return
		}
		if value >= 100 && value <= 2000 {
			p.height = value
		} else {
			http.Error(w, "height must be between 100 and 2000", http.StatusBadRequest)
			return
		}
	}
	if s := r.URL.Query().Get("top"); s != "" {
		top, err := parseColor(s)
		if err != nil {
			http.Error(w, "Error of parse top", http.StatusBadRequest)
			return
		}
		p.top = top
	}
	if s := r.URL.Query().Get("bottom"); s != "" {
		bottom, err := parseColor(s)
		if err != nil {
			http.Error(w, "Error of parse bottom", http.StatusBadRequest)
			return
		}
		p.bottom = bottom
	}

	p.xyscale = float64(p.width) / 2 / xyrange
	p.zscale = float64(p.height) * 0.4
	w.Header().Set("Content-Type", "image/svg+xml")
	drawingSVG(w, p)
}

func parseColor(s string) (color, error) {
	var c color
	if len(s) != 6 {
		return c, errors.New("error of input color")
	}
	var rgb [3]uint8
	for i := range 3 {
		iColor, err := strconv.ParseUint(s[2*i:2*i+2], 16, 8)
		if err != nil {
			return c, errors.New("error of parsing of color")
		}
		rgb[i] = uint8(iColor)
	}
	c.red, c.green, c.blue = rgb[0], rgb[1], rgb[2]
	return c, nil
}
