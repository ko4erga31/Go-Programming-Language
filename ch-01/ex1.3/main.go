/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.3
 *
 * Task:
 *   Measure the difference in running time between the potentially
 *   inefficient versions of echo and the one that uses strings.Join
 * =========================================================================
 */

package main

import (
	"fmt"
	"strings"
	"time"
)

func catWithJoin(words []string) string {
	var str string
	str = strings.Join(words, "")
	return str
}

func catWithPlus(words []string) string {
	var str string
	for _, word := range words {
		str += word
	}
	return str
}

func createSlice(num int) []string {
	var words []string
	for range num {
		words = append(words, "Join vs +=")
	}
	return words
}

func main() {
	slice := createSlice(10000)

	start := time.Now()
	catWithPlus(slice)
	finish1 := time.Since(start)

	start = time.Now()
	catWithJoin(slice)
	finish2 := time.Since(start)

	fmt.Println("+= time =", finish1)
	fmt.Println("Join time = ", finish2)
}
