/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.2
 *
 * Task:
 *   Modify the echo program to print the index and value
 *   of each of its arguments, one per line.
 * =========================================================================
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("idx\targ")
	for i, arg := range os.Args {
		fmt.Printf("%3d\t%s\n", i, arg)
	}

}
