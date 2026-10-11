/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.1
 *
 * Task:
 *   Modify the echo program to also print os.Args[0],
 *   the name of the command that invoked it.
 * =========================================================================
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	var line, space string
	for _, arg := range os.Args {
		line += space + arg
		space = " "
	}
	fmt.Println(line)
}
