package main
import "fmt"
/**
	常量等号右侧必须是常量或者是常量表达式
	常量表达式的函数必须是内置函数	
*/
const (
	text = "im vincent"
	length = len(text)
)
func main() {
	fmt.Println(text);
	fmt.Println(length);
}