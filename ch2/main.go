package main

import (
	"fmt"

	"github.com/robertding/The-Go-Programming-Language/ch2/popcount"
)

func main() {
	p1 := popcount.PopCount1(64)
	fmt.Println(p1)
	return
}
