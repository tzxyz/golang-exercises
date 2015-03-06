package main

import "fmt"

type bintree struct {
	key         int
	left, right *bintree
}

func NewBintree(k int, l, r *bintree) *bintree {
	return &bintree{k, l, r}
}

func SetChildNode(t, l, r *bintree) {
	t.left = l
	t.right = r
}

func Insert(t *bintree, z *bintree) {
	var y *bintree
	x := t
	for x != nil {
		y = x //y用来保存x的指针，x向下遍历
		if x.key > z.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	// z.parent = y
	if y == nil { //该树是空的，z直接作为root
		t = z
	}
	if z.key < y.key { //判断是在y左侧还是右侧
		y.left = z
	} else {
		y.right = z
	}
}

func InOrderTreeWalk(t *bintree) {
	if t.left != nil {
		InOrderTreeWalk(t.left)
	}
	// fmt.Printf("%d ", t)		//可以看出指针的情况
	fmt.Printf("%d ", t.key)
	if t.right != nil {
		InOrderTreeWalk(t.right)
	}
}

func InsertTest(root, t *bintree) {
	InOrderTreeWalk(root)
	fmt.Println()
	Insert(root, t)
	InOrderTreeWalk(root)
}

func main() {
	lchild := NewBintree(3, nil, nil)
	rchild := NewBintree(9, nil, nil)
	root := NewBintree(5, lchild, rchild)
	SetChildNode(root, lchild, rchild)		//就可以不要定义parent了
	InsertTest(root, NewBintree(4, nil, nil))
}
