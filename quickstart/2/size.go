package main
import "fmt"

/**
	_下划线表示忽略这个数
	kb，mb，gb可以这样表示
 */
const (
	_  = iota
	kb = 1 << (10 * iota)
	mb 
	gb 
)
func main() {
	fmt.Println(kb);
	fmt.Println(mb);
	fmt.Println(gb);
}