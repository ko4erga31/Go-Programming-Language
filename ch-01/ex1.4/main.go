/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.4
 *
 * Task:
 *   Modify dup2 to print the names of all files in which each
 *   duplicated line occurs.
 * =========================================================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filesByLine := make(map[string][]string)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesByLine)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
				continue
			}
			countLines(f, counts, filesByLine)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filesByLine[line], " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesByLine map[string][]string) {
	input := bufio.NewScanner(f)
	filename := f.Name()
	for input.Scan() {
		lineText := input.Text()
		counts[lineText]++
		files := filesByLine[lineText]
		if len(files) == 0 || files[len(files)-1] != filename {
			filesByLine[lineText] = append(files, filename)
		}
	}
}
