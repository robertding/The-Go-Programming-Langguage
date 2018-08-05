package main

import (
	"fmt"
	"time"

	"github.com/robertding/The-Go-Programming-Language/ch2/popcount"
)

func main() {
	p1 := popcount.PopCount1(63)
	p2 := popcount.PopCount1(63)
	p3 := popcount.PopCount1(66)
	fmt.Println(p1, p2, p3)
	t1 := time.Now()
	var i uint64
	var loops uint64 = 100000000
	for i = 0; i < loops; i++ {
		popcount.PopCount1(i)
	}
	t1Secs := time.Since(t1).Seconds()
	t2 := time.Now()
	for i = 0; i < loops; i++ {
		popcount.PopCount2(i)
	}
	t2Secs := time.Since(t2).Seconds()
	t3 := time.Now()
	for i = 0; i < loops; i++ {
		popcount.PopCount3(i)
	}
	t3Secs := time.Since(t3).Seconds()
	fmt.Printf("Seconds t1 %.4f, t2 %.4f, t3 %.4f\n", t1Secs, t2Secs, t3Secs)
	return
}
