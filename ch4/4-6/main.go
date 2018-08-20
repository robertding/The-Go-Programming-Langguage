package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "中国  的       弼马温  "
	fmt.Println("removeDup a", removeSpace([]rune(a)))
}

func removeSpace(s []rune) string {
	out := s[:0]
	var preIsSpace bool
	for _, x := range s {
		if !unicode.IsSpace(x) {
			out = append(out, x)
			preIsSpace = false
		} else {
			if !preIsSpace {
				out = append(out, rune(' '))
			}
			preIsSpace = true
		}
	}
	return string(out)
}
