package main
import (
	"fmt"
)
func main() {
	ch := make(chan int,1024)	//如果是有缓冲的channel，在channel的缓冲被填充完之前是不会阻塞的
	for {
		select {
			case ch<-1:
			case ch<-0:
		}
		// for i:= range ch {
		// 	fmt.Printf("Value=%d",i)
		// 	fmt.Println()		//这里为什么会死锁？ range不停的在读取ch中的数据，但是select那边在写入，所以死锁
		// }
		i := <-ch
		fmt.Printf("vvvvvv%d",i)
	}

}