package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := [...]int{9, 8, 7, 6}
	c := [...]int{0, 1}
	d := [...]int{1}
	reverse0(a[:])
	reverse0(b[:])
	reverse0(c[:])
	reverse0(d[:])
	fmt.Println("reverse0 a", a)
	fmt.Println("reverse0 b", b)
	fmt.Println("reverse0 c", c)
	fmt.Println("reverse0 d", d)
	a = [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println("reverse a", a)
}

func reverse0(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse(s *[6]int) {
	for i, j := 0, 6-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
