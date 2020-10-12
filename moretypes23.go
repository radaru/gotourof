package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	t := strings.Fields(s)
	ret :=  make(map[string]int)
	for _, element := range t {
		_, ok := ret[element]
		if ok {
			ret[element]++
		} else {
			ret[element] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
