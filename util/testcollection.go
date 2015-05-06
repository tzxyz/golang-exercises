package main

import (
	"fmt"
	"godemo/util/collection"
)

func main() {
	TestArrayList()
	// TestHashSet()
}

func TestArrayList() {
	// fmt.Println(collection.DEFAULT_EMPTY_SLICE)
	arrayList := collection.NewArrayList()
	// arrayList.Add("aaa")
	// fmt.Println(arrayList)
	fmt.Println(arrayList.Containes("aaa"))
	fmt.Println(arrayList.Size())
	fmt.Println(arrayList.IsEmpty())

	arrayList.Add("vvv")
	arrayList.Delete("vvv") // ERROR

	fmt.Println(arrayList)
	fmt.Println(arrayList.Size())
}

func TestHashSet() {
	hashset := collection.NewHashSet()

	hashset.Add("a")
	hashset.Add("a")
	hashset.Add("gg")

	fmt.Println(hashset)

	fmt.Printf("size is %d", hashset.Size())

	fmt.Printf("is empty? %v\n", hashset.IsEmpty())

	fmt.Printf("contains gg? %v\n", hashset.Contains("gg"))

	// hashset.Clear()
	// fmt.Println(hashset)

	hashset.Delete("gg")
	fmt.Println(hashset)
}

func TestLinkedList() {
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
