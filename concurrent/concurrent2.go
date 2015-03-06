package main
import (
	"fmt"
	"sync"
	"runtime"
)

var count int = 0	//在函数外面声明的话就是全局变量

func main() {
	lock := &sync.Mutex{}	//创建一个锁
	for i := 0; i < 100; i++ {
		go Count(lock)		
	}	

	for {
		lock.Lock()
		c := count
		fmt.Println(c)
		lock.Unlock()
		runtime.Gosched()
		if c>=10000 {
			break
		}
	}
}

func Count(lock *sync.Mutex) {
	lock.Lock()
	for i := 0; i < 100; i++ {
		count++
	}
	lock.Unlock()
}