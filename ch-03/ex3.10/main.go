/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.10
 *
 * Task:
 *   Write a non-recursive version of comma, using bytes.Buffer
 *   instead of string concatenation.
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

func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var b bytes.Buffer

	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i > 0 {
			b.WriteByte(',')
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

func main() {
	fmt.Println(comma1("123456789012002"))
}
