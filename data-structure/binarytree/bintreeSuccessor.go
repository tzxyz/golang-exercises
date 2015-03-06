package main

import "fmt"

type bintree struct {
	key                 int      //该节点上的数据
	parent, left, right *bintree //封装三个指针 parent：父节点，left：左孩子，right：右孩子
}

func NewBinTree(k int, l, r *bintree) *bintree {
	return &bintree{k, nil, l, r}
}

func Successor(r *bintree, val int) int { //找出在指定二叉树上，指定值的后继
	if val == r.Max() {
		return -1 //返回-1表示没有后继
	}
	t := Seacher(r, val)
	if t.right != nil {
		return t.right.Min()
	}
	p := t.parent
	for p != nil && p.right == t {
		t = p
		p = p.parent
	}
	return p.key
}

func Seacher(t *bintree, val int) *bintree {
	if t.key == val || t == nil {
		return t
	}
	if t.key > val {
		return Seacher(t.left, val)
	} else {
		return Seacher(t.right, val)
	}

}

func (t *bintree) SetParent(p *bintree) {
	t.parent = p
}

func (t *bintree) Min() int {
	for t.left != nil {
		t = t.left
	}
	return t.key
}

func (t *bintree) Max() int {
	for t.right != nil {
		t = t.right
	}
	return t.key
}

func main() {
	llchild := NewBinTree(2, nil, nil)
	lrchild := NewBinTree(4, nil, nil)
	rlchild := NewBinTree(6, nil, nil)
	rrchild := NewBinTree(10, nil, nil)
	lchild := NewBinTree(3, llchild, lrchild)
	rchild := NewBinTree(9, rlchild, rrchild)
	root := NewBinTree(5, lchild, rchild)

	llchild.SetParent(lchild)
	lrchild.SetParent(lchild)
	rlchild.SetParent(rchild)
	rrchild.SetParent(rchild)
	rchild.SetParent(root)
	lchild.SetParent(root)
	fmt.Println(Successor(root, 4))
	fmt.Println(Successor(root, 3))
	fmt.Println(Successor(root, 6))
	fmt.Println(Successor(root, 5))
	fmt.Println(Successor(root, 10))
}
