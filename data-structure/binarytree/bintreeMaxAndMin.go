package main
import "fmt"

type bintree struct {
	key int
	left, right *bintree
}

func (t *bintree) Min() int {
	for t.left !=nil {
		t = t.left
	}
	return t.key
}

func (t *bintree) Max() int {
	for t.right !=nil {
		t = t.right
	}
	return t.key
}

func main() {
	llchild := &bintree{2, nil, nil}
	lrchild := &bintree{4, nil, nil}
	rlchild := &bintree{6, nil, nil}
	rrchild := &bintree{10, nil, nil}
	lchild := &bintree{3, llchild, lrchild}
	rchild := &bintree{9, rlchild, rrchild}
	root := &bintree{5, lchild, rchild}
	max := root.Max()
	min := root.Min()
	fmt.Println("max = ", max)
	fmt.Println("min = ", min)
}