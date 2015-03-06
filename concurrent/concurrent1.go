package main
import (
	"fmt"
	"time"
)
func main() {
	for i := 0; i < 5; i++ {
		go add(i,i)
	}
	time.Sleep(3*time.Second)
	//main函数结束不会等待其他函数执行
}

func add(x,y int){
	z := x + y
	fmt.Println(z)
}