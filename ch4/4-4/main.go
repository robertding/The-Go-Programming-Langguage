package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:])
	fmt.Println("rotate a", a)
}

func rotate(s []int) {
	start := s[0]
	copy(s[:], s[1:])
	s[len(s)-1] = start
}
