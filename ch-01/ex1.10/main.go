/*
 * =========================================================================
 * Copyright (c) 2026 Illiryca
 * SPDX-License-Identifier: MIT
 *
 * Source:   The Go Programming Language, A. Donovan & B. Kernighan
 * Exercise: 1.10
 *
 * Task:
 *   Find a web site that produces a large amount of data. Investigate
 *   caching by running fetchall twice in succession to see whether the
 *   reported time changes much. Do you get the same content each time?
 *   Modify fetchall to print its output to a file so it can be examined.
 *
 * Run:
 *   go run main.go https://www.firefox.com https://google.com https://www.microsoft.com
 *
 * NB:
 *   Repeated runs are noticeably faster: DNS lookups and server/CDN
 *   caches are warm. Content of static pages is identical between
 *   runs, dynamic pages (e.g. google.com) differ slightly.
 * =========================================================================
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

	nameUrl := urlToFilename(url)
	file, err := os.Create(nameUrl)
	if err != nil {
		ch <- fmt.Sprintf("Error of create file %s: %v", nameUrl, err)
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func urlToFilename(url string) string {
	var clearName string

	clearName = strings.TrimPrefix(url, "https://")
	clearName = strings.TrimPrefix(clearName, "http://")
	clearName = strings.TrimPrefix(clearName, "www.")
	clearName = strings.TrimSuffix(clearName, ".com")
	clearName = strings.TrimSuffix(clearName, ".org")

	return clearName + ".html"
}
