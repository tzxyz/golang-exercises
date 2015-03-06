package main

import "fmt"

type Rect struct {
	width, height float64
}

	/**
	 *	结构体里不能有方法
	 *	必须显示this，即 func (r *Rect)
	 */
func (r *Rect) area(width float64, height float64) float64 {
	return r.width * r.height
}

func main() {
	/**
	 *	 自定义结构体都是引用类型 所以Rect前面必须加上*
	 */
	var r *Rect = &Rect{width: 1.5, height: 1.5}
	var area float64 = r.area(r.width,r.height)
	fmt.Println(area)
}
