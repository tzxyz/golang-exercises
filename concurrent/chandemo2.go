package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 可以看大片没有使用runtime.GOMAXPROCS()的时候
// 输出结果都是按照顺序的，说明只有一个系统进程服务与goroutine
// 该例子的作用是开启与cpu核数一样多的goroutine去计算累加到uint32的最大值
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(i int) {
			defer wg.Done()
			sum(i)
		}(i)
	}
	wg.Wait()
}

// 花费的时间也不一样
func sum(i int) {
	now := time.Now()
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	fmt.Println(i, x)
	fmt.Println("time ...", time.Now().Sub(now))
}
