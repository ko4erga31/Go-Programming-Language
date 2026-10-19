/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 3.13
 *
 * Task:
 *   Write const declarations for KB, MB, up through YB as compactly
 *   as you can.
 *
 * Run:
 *   go run .
 * =========================================================================
 */

package main

import "fmt"

const KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
	fmt.Println(YB / ZB)

	var b int64 = EB
	// var z int64 = ZB // error: 1e21 overflows int64
	fmt.Println(b)

}
