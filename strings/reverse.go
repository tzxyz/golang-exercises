package strings

// 给定一个字符串，要求把字符串前面的若干个字符移动到字符串的尾部
// 如把字符串“abcdef”前面的2个字符'a'和'b'移动到字符串的尾部
// 使得原字符串变成字符串“cdefab” 请写一个函数完成此功能
// 要求对长度为n的字符串操作的时间复杂度为 O(n)，空间复杂度为 O(1)。

// 暴力解法
func reverseViolent(raw string, to, length int) string {
	// 将string转换成字符数组
	count := 0
	runes := []rune(raw)
	for i := to; i > 0; i-- {
		// 每次都记住第一位
		tmp := runes[0]
		// 从第二位开始,直到to为止
		for j := 1; j < length; j++ {
			// 每一个元素向前移一位
			runes[j-1] = runes[j]
			count++
		}
		// 将原先第一位的值赋给最后一位
		runes[length-1] = tmp
	}
	return string(runes)
}

// 三步反转法
// 劳资tm的真的不会取名字了!
func reverseThree(raw string) {

}
