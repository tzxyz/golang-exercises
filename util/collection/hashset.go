package collection

import (
	"fmt"
)

type HashSet struct {
	hashmap map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{
		hashmap: make(map[interface{}]bool),
	}
}

func (this *HashSet) Size() int {
	return len(this.hashmap)
}

func (this *HashSet) IsEmpty() bool {
	return len(this.hashmap) == 0
}

func (this *HashSet) Contains(val interface{}) bool {
	_, ok := this.hashmap[val]
	return ok
}

func (this *HashSet) Add(val interface{}) {
	this.hashmap[val] = true
}

func (this *HashSet) Clear() {
	this.hashmap = make(map[interface{}]bool)
}

func (this *HashSet) Delete(val interface{}) {
	delete(this.hashmap, val)
}

func (this *HashSet) String() string {
	ret := "["
	for k, _ := range this.hashmap {
		ret += fmt.Sprintf("%v, ", k)
	}
	ret += "]"
	return ret
}
