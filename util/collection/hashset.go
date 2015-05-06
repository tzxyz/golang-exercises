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

func (h *HashSet) Size() int {
	return len(h.hashmap)
}

func (h *HashSet) IsEmpty() bool {
	return len(h.hashmap) == 0
}

func (h *HashSet) Contains(val interface{}) bool {
	_, ok := h.hashmap[val]
	return ok
}

func (h *HashSet) Add(val interface{}) {
	h.hashmap[val] = true
}

func (h *HashSet) Clear() {
	h.hashmap = make(map[interface{}]bool)
}

func (h *HashSet) Delete(val interface{}) {
	delete(h.hashmap, val)
}

func (h *HashSet) String() string {
	ret := "["
	for k, _ := range h.hashmap {
		ret += fmt.Sprintf("%v, ", k)
	}
	ret += "]"
	return ret
}
