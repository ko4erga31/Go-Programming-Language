/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 2.3
 *
 * Task:
 *   Rewrite PopCount to use a loop instead of a single expression.
 *   Compare the performance of the two versions.
 *
 * Run:
 *   go test -bench=.
 * =========================================================================
 */

package popcount

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountV1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountV2(x uint64) int {
	var num byte
	for i := range 8 {
		num += pc[byte(x>>(i*8))]
	}
	return int(num)
}
