package main

import (
	"fmt"
	"godemo/util/collection"
)

func main() {
	linkedList := collection.NewLinkedList()

	linkedList.Add("a")
	linkedList.Add("b")
	linkedList.Add("a")
	linkedList.Add("b")
	linkedList.Add("a")
	linkedList.Add("b")

	fmt.Println(linkedList.Size())
	fmt.Println(linkedList.IsEmpty())
	fmt.Println(linkedList.IndexOf("b"))
	fmt.Println(linkedList.Delete("a"))
	// fmt.Println(linkedList)
	// linkedList.Clear()
	fmt.Println(linkedList.Contains("b"))
	fmt.Println(linkedList)

	// ok := linkedList.(collection.Collection)
	// fmt.Println(ok)

	// fmt.Println(linkedList.(type))
}
