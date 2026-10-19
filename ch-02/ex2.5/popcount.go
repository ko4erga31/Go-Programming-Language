/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 2.5
 *
 * Task:
 *   Use the expression x&(x-1), which clears the rightmost non-zero bit,
 *   to write a faster version of PopCount. Compare the performance of
 *   the three versions.
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

func PopCountV4(x uint64) int {
	var count int
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
