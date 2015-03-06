package main

import "strings"
import "fmt"

func main() {
	var str1 string = "fuck my life"
	var str2 string = "fuck u"

	//strings.EqualFold(s, t string)：判断字符串s，t是否相同
	fmt.Println("1.判断字符串s，t是否相同",strings.EqualFold(str1, str2))

	var prefix string = "fuck"

	//strings.HasPrefix(s, prefix)：判断字符串s是否是以prefix开头
	fmt.Println("2.判断字符串s是否是以prefix开头",strings.HasPrefix(str1, prefix))

	var suffix string = "life"

	//strings.HasSuffix(s, suffix)：判断字符串s是否是以suffix结尾
	fmt.Println("3.判断字符串s是否是以suffix结尾",strings.HasSuffix(str1, suffix))

	var contains string = "my"

	//strings.Contains(s, c string)：判断字符串s是否包含字符串c
	fmt.Println("4.判断字符串s是否包含字符串c",strings.Contains(str1, contains))

	var r rune = 'f'

	//strings.ContainsRune(s string, c rune)：判断字符串s是否包含字符c
	fmt.Println("5.判断字符串s是否包含字符c",strings.ContainsRune(str1, r))

	//string.ContainsAny(s, any string)：判断字符串s是否包含字符串any中的任意字符
	fmt.Println("6.判断字符串s是否包含字符串any中的任意字符", strings.ContainsAny(str1, str2))

	//strings.Count(s, sep string)：判断字符串s中有多少个sep子字符串
	fmt.Println("7.判断字符串s中有多少个sep子字符串", strings.Count(str1, prefix))

	//strings.Index(s, sep string)：获取字符串s中第一次出现子字符串sep的位置
	fmt.Println("8.获取字符串s中第一次出现子字符串sep的位置", strings.Index("hello boss", "boss"))

	//strings.IndexByte(s string, b byte)：
	fmt.Println("")

	//strings.IndexRune(s string, r rune)：
	fmt.Println("")

	//strings.IndexAny(s, any string)：
	fmt.Println("")

	//string.LastIndex(s, sep string):
	fmt.Println()

	//strings.Title(s string)：字符串s中每个单词首字母大写
	fmt.Println("", strings.Title("fuck u son_of"))

	//strings.ToLower(s string)：返回字符串s所有字母改成小写的字符串
	fmt.Println("", strings.ToLower("FUCK MY LIFE"))

	//strings.ToUpper(s string)：返回字符串s所有字母改成大小的字符串
	fmt.Println("", strings.ToUpper("fuck my life"))

	//strings.Repeat(s string, count int)
	fmt.Println("", strings.Repeat("fuck ", 7))

	//strings.Replace(str,old, new string, n int)：将字符串str中的old替换成new，总共替换n次(n<0)全部替换
	fmt.Println("", strings.Replace("fuck fuck fuck my life", "fuck", "suck" , -1))

	f := func(r rune) rune{
		if r == 'm' || r == 'y' {
			return 'f'
		}
		return r
	}


	//strings.Map(mapping func(rune) rune, s string)：将字符串s分解为每个字符传入func，根据函数返回的rune替换
	fmt.Println("", strings.Map(f, "fuck my life"))

	//strings.Trim(s,trim string)：去除字符串s两端的字符串trim
	fmt.Println("", strings.Trim("***strings***","*"))

	//strings.TrimSpace(s string)：去除字符串s两端的空白
	fmt.Println("", strings.TrimSpace("   hahaha    "))

	f1 := func(r rune) bool{
		if r == 'h' {
			return true
		}
		return false
	}

	//strings.TrimFunc(s string, f func(rune) rune bool)：将字符串s分解为每个字符传入func，如果返回值为true就去除(两端)
	fmt.Println("", strings.TrimFunc("hhhxhhxhhxhhh", f1))

	//strings.Fields(s string) []string：按照空格将字符串分割成字符串数组
	fmt.Println("", strings.Fields("fuck my life"))

	f2 := func(r rune) bool {	//按照 | 和 $ 分割字符串
		if r == '|' || r == '$' {
			return true
		}
		return false
	}

	//strings.FieldsFunc(s string, f func(rune) bool) []string：将字符串按照指定函数进行分割(将字符串分解成每个字符传入函数，判断是否符合条件)
	fmt.Println("", strings.FieldsFunc("fuck$my|life", f2))

	//strings.Split(s, sep string) []string：将字符串s按照sep切分，如果sep为空，分割成每一个字符串
	fmt.Println("", strings.Split("fuck u", ""))

	//strings.SplitN(s, sep string, n int) []string：将字符串s按照sep切分，共分割成n份(n==0 返回nil)
	fmt.Println("", strings.SplitN("fuck u", "", 2))

	//strings.SplitAfter(s,sep string)：从出现sep之后按sep切分
	fmt.Println("", strings.SplitAfter("fuck u u u u ", "u"))

	//strings.Join(s []string, sep string) string：将字符串数组按照按照sep连接起来
	fmt.Println("", strings.Join([]string{"fuck","u"}, " "))

	
}
