package main

import "golang.org/x/tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, value := range strings.Fields(s) {
		m[value]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
