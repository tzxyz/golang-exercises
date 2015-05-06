package collection

import "fmt"

type Collection interface {
	Size() int
	IsEmpty() bool
	Contains(interface{}) bool
	Add()
	Delete(interface{}) interface{}
	Clear()
}

type List interface {
	Collection
	IndexOf(i interface{}) int
	Set(index int, i interface{}) interface{}
	Get(index int) interface{}
}

type Set interface {
	Collection
}

// linkedlist结构
type LinkedList struct {
	first, last *Node
	size        int
}

// 节点
type Node struct {
	item       interface{}
	prev, next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(val interface{}) {
	// 新节点
	newNode := &Node{
		item: val,
	}
	// 链表为空
	if l.last == nil {
		l.first = newNode
		l.last = newNode
	} else {
		l.last.next = newNode
		l.last = newNode
	}

	// 容量+1
	l.size++
}

// 删除节点
func (l *LinkedList) Delete(val interface{}) interface{} {
	for x := l.first; x != nil; x = x.next {
		if x.item == val {
			prev := x.prev
			next := x.next

			if prev == nil {
				l.first = x.next
			} else {
				prev.next = x.next
				x.item = nil
			}

			if next == nil {
				l.last = x.prev
			} else {
				next.prev = x.prev
				x.item = nil
			}

			l.size-- // 容量-1

			return val
		}
	}
	return nil
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) String() string {
	if l.size == 0 {
		return "[]"
	}
	ret := ""

	for x := l.first; x != nil; x = x.next {
		ret += fmt.Sprint("", x.item)
	}

	return ret
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) IndexOf(val interface{}) int {
	index := 0
	for x := l.first; x != nil; x = x.next {
		if x.item == val {
			return index
		}
		index++
	}
	return -1
}

// 清空list
func (l *LinkedList) Clear() {
	for x := l.first; x != nil; x = x.next {
		x.item = nil
		x.prev = nil
		x.next = nil
	}
	l.first = nil
	l.last = nil
	l.size = 0
}

// 获取
func (l *LinkedList) Get(index int) interface{} {
	ok := l.checkIndex(index)
	if !ok {

	}
	tmp := 0
	for x := l.first; x != nil; x = x.next {
		tmp++
		if tmp == index {
			return x.item
		}
	}
	return nil
}

// 索引是否合法
func (l *LinkedList) checkIndex(index int) bool {
	if index >= 0 && index < l.size {
		return true
	}
	return false
}

// func (l *LinkedList) getNode(val interface{}) {

// }

func (l *LinkedList) Contains(val interface{}) bool {
	return l.IndexOf(val) != -1
}
