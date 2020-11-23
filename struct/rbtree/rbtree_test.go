package rbtree

import (
	"fmt"
	"testing"
)

func TestRBTree(t *testing.T) {
	tree := new(RBTree)
	tree.Insert(7).
		Insert(3).
		Insert(4).
		Insert(1).
		Insert(9).
		Insert(10).
		Insert(8)
	fmt.Println(tree)

}
