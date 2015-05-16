package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Key struct {
	prop   string `xml:"property"`
	column string `xml:"column"`
}

type Result struct {
	prop   string `xml:"property,attr"`
	column string `xml:"column,attr"`
}

type ResultMap struct {
	XMLName xml.Name `xml:"resultMap"`
	Id      string   `xml:"id,attr"`
	Type    string   `xml:"type,attr"`
	Key
	Result
}

// 前缀
const suffix = "ResultMap"

// 期望缩进
const indent = "	"

var attributes []string = make([]string, 0, 1024)

func main() {
	// 1.读取DTO文件
	file, err := os.Open("User.java")

	if err != nil {
		fmt.Println("a")
	}

	defer file.Close()
	// 2.读取包名，类名（第一行就是包，类名是文件名），type就解决了，id用类名+ResultMap
	className := strings.Split(file.Name(), ".")[0] // file.Name()是拿到全名

	id := className + suffix // fix id

	fmt.Println(id)

	buffer := bytes.NewBuffer(make([]byte, 0, 1024)) // 缓存

	bufferdata, err := ioutil.ReadFile("User.java")
	buffer = bytes.NewBuffer(bufferdata)
	if err != nil {
		fmt.Println("aa")
	}

	reader := bufio.NewReader(buffer) // 读取流

	packageName, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println("readline err:", err)
	}
	data, _, err := reader.ReadLine()
	fmt.Println("data", data, "err", err)

	// 读取剩余的
	// for line, _, err := reader.ReadLine(); err == nil; {
	// 	// fmt.Println("LINE:", string(line))
	// 	s := string(line)
	// 	fmt.Println("s", s)
	// }

	for i := 1; ; {
		line, err := reader.ReadString('\n') //我日
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println("linenum:", i, line)
		i += 1
	}

	packageName = bytes.Split(packageName, []byte(" "))[1]
	packageName = packageName[:len(packageName)-1]

	typeName := string(packageName) + "." + className
	// fmt.Println(typeName)
	// fmt.Println(string(packageName))

	os.Stdout.Write(packageName)

	resultMap := &ResultMap{
		Id:   id,
		Type: typeName,
	}

	data, err = xml.MarshalIndent(resultMap, indent, indent)
	if err != nil {
		fmt.Println("xml marshal err")
	}
	// fmt.Println(resultMap)
	// fmt.Println(string(data))
}
