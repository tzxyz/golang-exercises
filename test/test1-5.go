package main

import "fmt"

func main() {
	f:= []float64{1.1,2.2,3.3,4.4,5.5}
	fmt.Printf("平均数为：%f",Avg(f))
}

func Avg(f []float64) float64 {
	var sum float64
	length := len(f)
	for i := 0 ; i < length; i++ {
		sum += f[i]
	}
	avg := sum/float64(length)	//int类型不能和float类型做计算
	return avg
}
