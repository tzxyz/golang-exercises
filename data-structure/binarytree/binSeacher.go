package main
import "fmt"
type bintree struct {
	key int 
	left, right *bintree
}

func TreeSeacher(x *bintree, val int) *bintree {
	if x.key == val {
		return x
	}
	if x.key < val {
		return TreeSeacher(x.right, val)
	} else {
		return TreeSeacher(x.left, val)
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

	tree := TreeSeacher(root, 9)
	fmt.Println(tree.key, tree.left, tree.right)
}