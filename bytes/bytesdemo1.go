package main

import (
	"bytes"
	"fmt"
)

func main() {
	// CompareDeme()
	// ContainsDemo()
	// CountDemo()
	// EqualDemo()
	// EqualFoldDemo()
	// FieldsDemo()
	// FieldsFuncDemo()
	// HasPrefixDemo()
	HasSuffixDemo()
}

// 9.bytes.HasSuffix(s, suffix) bool
// 判断切片的后缀是否为prefix
func HasSuffixDemo() {
	b := []byte("abcfefg")
	p1 := []byte("abc")
	p2 := []byte("efg")
	fmt.Println(bytes.HasSuffix(b, p1)) //Outout:false
	fmt.Println(bytes.HasSuffix(b, p2)) //Output:true
}

// 8.bytes.HasPrefix(s, prefix) bool
// 判断切片的前缀是否为prefix
func HasPrefixDemo() {
	b := []byte("abcfefg")
	p1 := []byte("abc")
	p2 := []byte("efg")
	fmt.Println(bytes.HasPrefix(b, p1)) //Output:true
	fmt.Println(bytes.HasPrefix(b, p2)) //Output:false
}

// 7.bytes.FieldsFunc(s, func(r rune) bool)
// func(r rune) bool 返回true,就把这个字符用于分割
func FieldsFuncDemo() {
	b := []byte("嘿啊嘿 哦啊哦 哈啊哈 呵啊呵")
	_bytes := bytes.FieldsFunc(b, func(r rune) bool { //返回的是bool,传入的r是否相等
		return r == '啊' //很少用到rune,注意要用''单引号包围,才表示字符
	})
	for index, _b := range _bytes {
		fmt.Printf("index=%d,byte=%d,str=%s\n", index, _b, string(_b))
	}
}

// 6.bytes.Fields(s) [][]byte
// 将s按照一个或多个空字符串分割,返回字节切片的切片
// \t,\n这种算多个空格
func FieldsDemo() {
	b := []byte("啊啊啊 啊\t啊 啊\n啊 啊啊啊") //中文也能转 说明用的是UTF-8
	_bytes := bytes.Fields(b)
	for index, _b := range _bytes {
		fmt.Printf("index=%d, byte=%d, str=%s\n", index, _b, string(_b))
	}
}

// 5.bytes.EqualFold(s,t) bool
// 将s,t全部转换成utf-8编码,并忽略大小写编码
// question:go怎么使用其他编码
func EqualFoldDemo() {
	b1 := []byte("ABC")
	b2 := []byte("abc")
	fmt.Println(bytes.EqualFold(b1, b2)) //Output:true
}

// 4.bytes.Equal(a,b) bool
// 判断两个字节切片是否相等
// ps:空切片和nil相等
func EqualDemo() {
	b1 := []byte("fuck my life")
	b2 := []byte("fuck my 1ife")
	b3 := []byte{}
	fmt.Println(bytes.Equal(b1, b2))  //Output:false
	fmt.Println(bytes.Equal(b3, nil)) //Output:true

	// fmt.Println(b3, nil) 有意思的现象
}

// 3.bytes.Count(s,sep) int
// 子切片出现的次数
func CountDemo() {
	b1 := []byte("abcdffffffffffuck")
	b2 := []byte("abc")
	b3 := []byte("f")
	b4 := []byte("ff")

	fmt.Println(bytes.Count(b1, b2)) //Output:1
	fmt.Println(bytes.Count(b1, b3)) //Output:10
	fmt.Println(bytes.Count(b1, b4)) //Output:5
}

// 2.bytes.Contains(s,subslice) bool
// 判断切片是否存在包含关系
func ContainsDemo() {
	b1 := []byte("abc")
	b2 := []byte("ab")
	b3 := []byte("bd")
	fmt.Println(bytes.Contains(b1, b2)) // Output:true
	fmt.Println(bytes.Contains(b1, b3)) // Output:false
}

// 1.bytes.Compare(a,b) int
// 比较两个字节切片的大小
// 大返回1,小返回-1,一样大返回0
func CompareDeme() {
	b1 := []byte{0}
	b2 := []byte{2}

	// b3 := []byte{"a"} 字节数组只能0-255的数字 2^8-1
	// b4 := []byte{"c"} 1byte = 8bit

	b3 := []byte("a")
	b4 := []byte("b")

	b5 := []byte{255}
	// b5 := []byte{256}

	fmt.Println(bytes.Compare(b1, b2)) // Output:-1
	fmt.Println(bytes.Compare(b3, b4)) // Output:-1
	fmt.Println(b5)
}
