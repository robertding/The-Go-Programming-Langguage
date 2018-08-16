package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abca"
	b := "abc"
	c := "aabd"
	fmt.Println(a, b, sameSentenceIsomerism(a, b))
	fmt.Println(a, c, sameSentenceIsomerism(a, c))
	fmt.Println(c, b, sameSentenceIsomerism(c, b))
}

func sameSentenceIsomerism(a, b string) bool {
	s1 := strings.Trim(a, b) == ""
	s2 := strings.Trim(b, a) == ""
	return s1 && s2
}
