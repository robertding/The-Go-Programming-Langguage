package main

import (
	"fmt"
)

func main() {
	a := [...]string{"a", "b", "a", "a", "b"}
	fmt.Println("removeDup a", removeDup(a[:]))
}

func removeDup(s []string) []string {
	out := s[:0]
	var pre string
	for _, x := range s {
		if x != pre {
			out = append(out, x)
			pre = x
		}
	}
	return out
}
