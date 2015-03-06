package main

import (
	"fmt"
	"regexp"
)

func main() {
	// ValidID()
	// MatchString()
	// UsuallyPhoneRegexp()
	UsuallyEmailPattern()
}

func ValidID() {
	// 编译表达式
	// ^表示开头 $表示结尾
	// [0-9]表示可选的字符 {2,5}可选字符的长度在2~5之间
	valid := regexp.MustCompile("^[0-9]{2,5}$")
	println(valid.MatchString("33"))
	println(valid.MatchString("a33"))
	println(valid.MatchString("999999"))
}

func MatchString() {
	// 使用正则匹配字符串
	pattern := "^0[0-9]{2,3}-[0-9]{5}$"
	matched, err := regexp.MatchString(pattern, "010-55555")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(pattern, "020-8888")
	fmt.Println(matched, err)

	// err返回的是编写错误的正则
	pattern = "a(b"
	matched, err = regexp.MatchString(pattern, "aaab")
	fmt.Println(matched, err)
}

func UsuallyPhoneRegexp() {
	// 手机号码的正则表达式
	pattern := "^1[3|5|8][0-9]{9}$"
	phones := []string{"15859009096", "18511669684", "18650467117", "17211229999"}
	for _, phone := range phones {
		if matched, _ := regexp.MatchString(pattern, phone); matched {
			fmt.Println(phone, "是正确的手机号码")
		} else {
			fmt.Println(phone, "不是正确的手机号码")
		}
	}
}

func UsuallyEmailPattern() {
	// ?表示匹配0次或多次
	// +表示一次或多次
	// \d表示匹配数组
	// \w表示匹配字母,汉字或下划线
	// .表示匹配除换行符以外的所有字符
	// [^x]表示除了x以外的所有字符
	pattern := `^([a-zA-Z0-9]+[_\_\.]?)@([a-zA-Z0-9]+)\.[a-zA-Z]{2,3}$`
	emails := []string{"876553848@qq.com", "kevinmesss.tz@gmail.com", "vincent.fml@hotmail.com"}
	for _, email := range emails {
		if matched, _ := regexp.MatchString(pattern, email); matched {
			fmt.Println(email, "格式正确")
		} else {
			fmt.Println(email, "格式不正确")
		}
	}
}
