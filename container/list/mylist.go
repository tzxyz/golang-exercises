package main

import (
	"container/list"
	"fmt"
)

func main() {
	// ListNewDemo()
	// ListFrontDemo()
	// ListBackDemo()
	// ListPushFront()
	// ListPushBack()
	// ListInit()
	// ListRemove()
	// ListPrev()
	// ListNext()
	// ListInsertBefore()
	// ListInsertAfter()
	ListMove2Front()
	ListMove2Back()
}

func ListInsertBefore() {
	l := list.New()
	l.PushBack("aaaaaa")
	l.PushBack("bbbbbb")
	l.PushBack("cccccc")
	l.InsertBefore("fuck", l.Front()) //添加在指定节点之前
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func ListInsertAfter() {
	l := list.New()
	l.PushBack("aaaaaa")
	l.PushBack("bbbbbb")
	l.PushBack("cccccc")
	l.InsertAfter("fuck", l.Back()) //添加在指定节点之后
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func ListMove2Back() {
	l := list.New()
	l.PushBack("aaaaaa")
	l.PushBack("bbbbbb")
	l.PushBack("cccccc")
	l.MoveToBack(l.Front()) //将指定节点移动到尾节点
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func ListMove2Front() {
	l := list.New()
	l.PushBack("aaaaaa")
	l.PushBack("bbbbbb")
	l.PushBack("cccccc")
	l.MoveToFront(l.Back()) //将指定节点移动到头节点
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func ListNext() {
	l := list.New()
	l.PushBack("aaaaaa")
	l.PushBack("bbbbbb")
	l.PushBack("cccccc")
	for e := l.Front(); e != nil; e = e.Next() { //获取e的后一个元素指针
		fmt.Println(e.Value)
	}
}

func ListPrev() {
	l := list.New()
	l.PushFront("aaaaa")
	l.PushFront("bbbbb")
	l.PushFront("ccccc")
	for e := l.Back(); e != nil; e = e.Prev() { //获取前一个元素的指针
		fmt.Println(e.Value)
	}
}

func ListRemove() {
	l := list.New()
	l.PushFront("a")
	fmt.Printf("remove befor ... len = %d\n", l.Len())
	e := l.Front()
	l.Remove(e) //删除链表的元素,e必须是这个链表的元素,而且类型是*list.Element
	fmt.Printf("remove after ... len = %d\n", l.Len())
}

func ListInit() {
	l := list.New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("c")
	fmt.Printf("list init befor ... len = %d\n", l.Len())
	l.Init() //清空/初始化链表 初始化的链表长度为0
	fmt.Printf("list init after ... len = %d\n", l.Len())
}

func ListPushBack() {
	l := list.New()
	l.PushBack("a") //从尾指针的方向向链表压入元素
	l.PushBack("b")
	l.PushBack("c")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func ListPushFront() {
	l := list.New()
	l.PushFront("a") //从头指针的方向向链表压入元素
	l.PushFront("b")
	l.PushFront("c")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
xx
func ListBackDemo() {
	l := list.New()
	e := l.Back() //获取list的尾指针
	if e == nil {
		l.PushBack("a")
	}
	fmt.Println(l.Back().Value) //获取元素的值
}

func ListFrontDemo() {
	l := list.New()
	e := l.Front() //获取list的头指针
	if e == nil {
		l.PushFront("a")
	}
	fmt.Println(l.Front().Value) //获取元素的值
}

func ListNewDemo() {
	l := list.New() //创建一个空链表 长度为0 头/尾指针为nil
	e := l.Front()
	fmt.Println(l)
	fmt.Println(e)
}
