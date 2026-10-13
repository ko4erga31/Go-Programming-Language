/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 2.2
 *
 * Task:
 *   Write a general-purpose unit-conversion program analogous to cf
 *   that reads numbers from its command-line arguments or from the
 *   standard input if there are no arguments, and converts each number
 *   into units like temperature in Celsius and Fahrenheit, length in
 *   feet and meters, weight in pounds and kilograms, and the like.
 *
 * Run:
 *   go run . 100 36.6
 *   go run . -length -weight 100
 *   echo "0 32" | go run . -temperature
 *
 * NB:
 *   Flags -temperature, -length, -weight select conversions; with no
 *   flags all of them are shown. Flags must precede the numbers.
 * =========================================================================
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ko4erga31/Go-Programming-Language/ch-02/ex2.2/lengthconv"
	"github.com/ko4erga31/Go-Programming-Language/ch-02/ex2.2/tempconv"
	"github.com/ko4erga31/Go-Programming-Language/ch-02/ex2.2/weightconv"
)

var l = flag.Bool("length", false, "use length convert")
var w = flag.Bool("weight", false, "use weight convert")
var t = flag.Bool("temperature", false, "use temperature")

func main() {
	flag.Parse()
	all := !*t && !*w && !*l
	showT, showW, showL := *t || all, *w || all, *l || all
	if len(flag.Args()) > 0 {
		for _, line := range flag.Args() {
			convert(line, showT, showW, showL)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			numbersInStr := strings.Fields(line)
			for _, numInStr := range numbersInStr {
				convert(numInStr, showT, showW, showL)
			}
		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
		}
	}
}

func convert(s string, t, w, l bool) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error of reading num: %v\n", err)
		return
	}
	fmt.Printf("Metrics for %s\n", s)
	if t {
		f, c := tempconv.Fahrenheit(num), tempconv.Celsius(num)
		fmt.Printf("%s = %s\t%s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	if w {
		p, k := weightconv.Pounds(num), weightconv.Kilograms(num)
		fmt.Printf("%s = %s\t%s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
	}
	if l {
		m, fe := lengthconv.Meters(num), lengthconv.Feet(num)
		fmt.Printf("%s = %s\t%s = %s\n", m, lengthconv.MToF(m), fe, lengthconv.FToM(fe))
	}
}
