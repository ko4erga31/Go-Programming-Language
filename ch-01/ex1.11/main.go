/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.11
 *
 * Task:
 *   Try fetchall with longer argument lists, such as samples from
 *   a list of the top million web sites. How does the program behave
 *   if a web site just doesn't respond?
 *
 * Run:
 *   go run main.go $(cat urls.txt)
 *
 * NB:
 *   Response bodies are discarded: only timing and size are reported.
 *   Unreachable hosts (e.g. 10.255.255.1) block the whole program:
 *   main waits on the channel until the slow goroutine finally times
 *   out, because http.Get has no timeout by default.
 * =========================================================================
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
