package main
import "fmt"
type bintree struct {
	key int 
	left, right *bintree
}

func PostOrderTreeWalk(x *bintree) {
	if x != nil {
		if x.left != nil {
			PostOrderTreeWalk(x.left)
		}
		if x.right != nil {
			PostOrderTreeWalk(x.right)
		}
		fmt.Println(x.key)
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
	PostOrderTreeWalk(root)
}