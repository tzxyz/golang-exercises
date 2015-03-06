package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	println(EncryptMd5("FUCK"))
	println(EncryptSha("FUCK"))
}

func EncryptMd5(raw string) string {
	m := md5.New()
	io.WriteString(m, raw)
	str := fmt.Sprintf("%x", m.Sum(nil))
	return str
}

func EncryptSha(raw string) string {
	s := sha1.New()
	io.WriteString(s, raw)
	str := fmt.Sprintf("%x", raw)
	return str
}
