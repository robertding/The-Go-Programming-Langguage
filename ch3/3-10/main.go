package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("123456789"))
}

func comma(s string) string {
	var buf bytes.Buffer
	lenS := len(s)
	mod3 := lenS % 3
	buf.WriteString(s[:mod3])
	for n := 0; n <= lenS-3; n += 3 {
		c := ","
		if n == 0 && mod3 == 0 {
			c = ""
		}
		buf.WriteString(c + s[mod3+n:mod3+n+3])
	}
	return buf.String()
}