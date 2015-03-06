package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	c := count("fuck my life")
	fmt.Println(c)
}

func count(s string) int{
	c := utf8.RuneCountInString(s)	//返回一个字符串的字节数
	b := utf8.
	fmt.Println(c)
	return c
}

func replace() {
	old := strings.
}
