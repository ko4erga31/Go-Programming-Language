/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.12
 *
 * Task:
 *   Modify the Lissajous server to read parameter values from the URL.
 *   For example, a URL like http://localhost:8000/lissajous?cycles=20
 *   sets the number of cycles to 20 instead of the default 5.
 *   Use strconv.Atoi to convert the string parameter into an integer.
 *
 * Run:
 *   go run .
 *   then open http://localhost:8000/lissajous?cycles=20
 *
 * NB:
 *   The cycles parameter is validated: values outside [1, 100] or
 *   non-numeric input fall back to the default. This prevents a client
 *   from making the server burn CPU on huge values (simple DoS).
 * =========================================================================
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", gifLis)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

const (
	defaultCycles = 5
	minCycles     = 1
	maxCycles     = 100
)

func gifLis(w http.ResponseWriter, r *http.Request) {
	cycles := defaultCycles

	if s := r.URL.Query().Get("cycles"); s != "" {
		n, err := strconv.Atoi(s)
		switch {
		case err != nil:
			log.Printf("invalid cycles %q: %v, using default %d", s, err, defaultCycles)
		case n < minCycles || n > maxCycles:
			log.Printf("cycles %d out of range [%d, %d], using default %d", n, minCycles, maxCycles, defaultCycles)
		default:
			cycles = n
		}
	}
	lissajous(w, cycles)
}
