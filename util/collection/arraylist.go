package collection

import (
	"fmt"
	"strings"
)

// 这样是个单例，应该不能这样
// var DEFAULT_EMPTY_SLICE []interface{}

const DEFAULT_ARRAY_SIZE = 10

type ArrayList struct {
	elements []interface{}
	size     int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elements: make([]interface{}, DEFAULT_ARRAY_SIZE),
		size:     0,
	}
}

func (this *ArrayList) Size() int {
	return this.size
}

func (this *ArrayList) IsEmpty() bool {
	return this.size == 0
}

func (this *ArrayList) Containes(val interface{}) bool {
	return this.IndexOf(val) != -1
}

// TODO 不能使用append()，slice默认有零值
func (this *ArrayList) Add(val interface{}) {
	this.elements = append(this.elements, val)
	this.size++
}

// 清空
func (this *ArrayList) Clear() {
	this.elements = make([]interface{}, DEFAULT_ARRAY_SIZE)
	this.size = 0
}

// TODO
// 删除切片特定的值，貌似只能设置为nil？
func (this *ArrayList) Delete(val interface{}) interface{} {
	// delete(this.elements, val)
	for i := 0; i < this.size; i++ {
		if this.elements[i] == val {
			// if strings.EqualFold(this.elements[i], val) {
			fmt.Println("delete")
			element := this.elements[i]
			this.elements[i] = nil
			this.size--
			return element
		}
	}
	fmt.Println("ddd")
	return nil
}

// 查找，找到返回元素第一次出现的索引，没找到返回-1
func (this *ArrayList) IndexOf(val interface{}) int {
	for index, data := range this.elements {
		if data == val {
			return index
		}
	}
	return -1
}

func (this *ArrayList) String() string {
	ret := "["
	for _, val := range this.elements {
		ret += fmt.Sprintf("%v ,", val)
	}
	ret += "]"
	return ret
}
