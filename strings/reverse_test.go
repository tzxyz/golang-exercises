package strings

import (
	"testing"
)

// 测试对象
// 使用了匿名结构体,注意要加逗号
var testDate = []struct {
	args     []interface{}
	expected string
}{
	{
		[]interface{}{"abcdefg", 3, len("abcdefg")},
		"defgabc",
	}, {
		[]interface{}{"sssmmm", 2, len("sssmmm")},
		"smmmss",
	}, {
		[]interface{}{"fffffuck", 5, len("fffffuck")},
		"uckfffff",
	},
}

// 测试ReverseViolent函数的正确性
func TestReverseViolent(t *testing.T) {
	for _, td := range testDate {
		result := reverseViolent(td.args[0].(string), td.args[1].(int), td.args[2].(int))
		if result != td.expected {
			t.Errorf("string reverse(%v) expected want(%s) but result(%s)", td.args, td.expected, result)
		}
	}
}
