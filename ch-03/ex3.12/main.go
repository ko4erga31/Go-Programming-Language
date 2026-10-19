/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.12
 *
 * Task:
 *   Write a function that reports whether two strings are anagrams
 *   of each other, that is, they contain the same letters in a
 *   different order.
 *
 * Run:
 *   go run .
 * =========================================================================
 */

package main

import (
	"fmt"
	"maps"
)

func isAnagramm(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return false
	}
	alphs1, alphs2 := make(map[rune]int), make(map[rune]int)

	for _, r := range s1 {
		alphs1[r]++
	}
	for _, r := range s2 {
		alphs2[r]++
	}

	return maps.Equal(alphs1, alphs2)
}

func main() {
	fmt.Println(isAnagramm("listen", "silent")) // true
	fmt.Println(isAnagramm("evil", "vile"))     // true
	fmt.Println(isAnagramm("dusty", "study"))   // true
	fmt.Println(isAnagramm("abc", "abd"))       // false
	fmt.Println(isAnagramm("aab", "abb"))       // false
	fmt.Println(isAnagramm("abc", "abcd"))      // false
	fmt.Println(isAnagramm("", ""))             // true
}
