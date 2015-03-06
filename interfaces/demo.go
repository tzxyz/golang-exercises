package main
import "fmt"
import "strconv"

type Element interface{}	
type List []Element	//切片

type Person struct {
	name string
	age int
}

func (p *Person) String() string{				//复写了fmt中的String方法都可以fmt中打印出字符串
	return "<" + p.name +"> <" + strconv.Itoa(p.age) + ">"
}

func main() {
	list := make(List, 3)	//make只能创建slice，map，channel，返回的是T类型不是*T类型
	list[0] = 0				//new返回的是指针*T类型
	list[1] = "fuck"
	list[2] = Person{"卓尼玛", 22}

	for i, e := range list {
		switch val := e.(type) {
		case int:
			fmt.Printf("list[%d] is an int type, val is %d\n", i, val)
		case string:
			fmt.Printf("list[%d] is a string type, val is %s\n", i, val)
		case Person:
			fmt.Printf("list[%d] is a Person type, val is %s\n", i, val)
		default:
			fmt.Printf("list[%d] is a diff type", i)
		}
	}
}

