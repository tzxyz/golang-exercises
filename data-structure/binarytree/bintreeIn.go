package main

import "fmt"

type bintree struct {
	key   int
	left  *bintree
	right *bintree
}

/**
 *	二叉树中序遍历：先输出T.root的左子节点，再输出右子节点
 */
func InOrderTreeWalk(x *bintree) {
	if x != nil {
		if x.left != nil {
			InOrderTreeWalk(x.left)
		}
		fmt.Println(x.key)
		if x.right != nil {
			InOrderTreeWalk(x.right)
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
	InOrderTreeWalk(root)
}
