package main

import "fmt"

type bintree struct {
	key         int
	left, right *bintree
}

/**
 *	先序遍历-递归
 */
func PreOrderTreeWalk(x *bintree) {
	if x != nil {
		fmt.Println(x.key)
		if x.left != nil {
			PreOrderTreeWalk(x.left)
		}
		if x.right != nil {
			PreOrderTreeWalk(x.right)
		}
	}
}

func main() {
	llchild := &bintree{2, nil, nil}
	lrchild := &bintree{4, nil, nil}
	rlchild := &bintree{6, nil, nil}
	rrchild := &bintree{10, nil, nil}
	lchild := &bintree{3, llchild, lrchild}
	rchild := &bintree{9, rlchild, rrchild}
	root := &bintree{5, lchild, rchild}
	PreOrderTreeWalk(root)
}
