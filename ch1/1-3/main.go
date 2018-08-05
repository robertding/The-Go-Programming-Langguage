package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs1 := time.Since(t1).Seconds()
	t2 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	secs2 := time.Since(t2).Seconds()
	fmt.Printf("t1 %.6f, t2 %.6f, diff %.6f\n", secs1, secs2, secs2-secs1)
}
