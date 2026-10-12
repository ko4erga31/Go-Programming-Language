/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.7
 *
 * Task:
 *   Use io.Copy instead of ioutil.ReadAll to copy the response body
 *   to os.Stdout without allocating a buffer large enough to hold
 *   the entire response. Check the error returned by io.Copy.
 *
 * Run:
 *   go run main.go https://example.com
 * =========================================================================
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy of %s: %v\n", url, err)
		}
	}
}
