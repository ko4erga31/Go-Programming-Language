/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.11
 *
 * Task:
 *   Enhance comma so that it deals correctly with floating-point
 *   numbers and an optional sign.
 *
 * Run:
 *   go run .
 * =========================================================================
 */

package main

import (
	"bytes"
	"fmt"
)

func checkFloat(s string, n int) (int, bool) {
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			return i, true
		}
	}
	return 0, false
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var b bytes.Buffer
	var isSign bool
	var start, end int
	idx, isFloat := checkFloat(s, n)
	if s[0] == '-' || s[0] == '+' {
		isSign = true
	}
	if isSign {
		start = 1
	}
	if isFloat {
		end = idx
	} else {
		end = n
	}

	for i := 0; i < n; i++ {
		if (end-i)%3 == 0 && i > start && end > i {
			b.WriteByte(',')
		}
		b.WriteByte(s[i])
	}

	return b.String()
}

func main() {
	fmt.Println(comma2("+10010923812.120312"))
	fmt.Println(comma2("1234567"))
	fmt.Println(comma2("-1234567"))
	fmt.Println(comma2("1234567.891"))
	fmt.Println(comma2("-1234567.891"))
	fmt.Println(comma2("+1234"))
	fmt.Println(comma2("123456"))
	fmt.Println(comma2("-123"))
	fmt.Println(comma2("12"))
	fmt.Println(comma2(".5"))
	fmt.Println(comma2("-0.123456"))
}
