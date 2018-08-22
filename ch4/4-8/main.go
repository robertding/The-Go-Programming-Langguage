package main

import (
	"fmt"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	str := "fie色粉fe 绿色积极��想哭��i2332323"
	for _, r := range str {
		if unicode.IsLetter(r) {
			counts[r]++
		}
	}
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
