/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.9
 *
 * Task:
 *   Modify fetch to also print the HTTP status code, found in
 *   resp.Status.
 *
 * NB:
 *   (This version defaults to https:// as a safer choice.)
 *
 * Run:
 *   go run main.go https://example.com http://example.com example.com
 * =========================================================================
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\tStatus=%d\n", url, resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy of %s: %v\n", url, err)
		}
	}
}
