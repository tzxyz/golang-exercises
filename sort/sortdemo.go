package main

import (
	"fmt"
	"reflect"
	"sort"
)

// OrderedMap
type OrderedMap struct {
	keys []interface{}
	dict map[interface{}]interface{}
}

// 实现sort包的三个接口
func (this *OrderedMap) Len() int {
	return len(this.keys)
}

func (this *OrderedMap) Less(i, j int) bool {
	fmt.Println("...")
	return true
}

func (this *OrderedMap) Swap(i, j int) {
	this.keys[i], this.keys[j] = this.keys[j], this.keys[i]
}

type KeyInterface interface {
	sort.Interface
	Add()
	Remove()
	Get(index int)
	Search(key interface{}) (index int, contains bool)
	Type() reflect.Type
	Compare() func(this, that KeyInterface) int8
}

type Key struct {
	container []interface{}
	keyType   reflect.Type
	compare   func(this, that interface{}) int8 // 为什么这里要用interface
}

func (this *Key) Len() int {
	return len(this.container)
}

func (this *Key) Less(i, j int) bool {
	return this.compare(this.container[i], this.container[j]) == -1 // TODO
}

func (this *Key) Swap(i, j int) {
	this.container[i], this.container[j] = this.container[j], this.container[i]
}
