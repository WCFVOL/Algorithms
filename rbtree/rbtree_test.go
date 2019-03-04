package rbtree

import (
	"fmt"
	"testing"
)

func less(k1, k2 interface{}) bool {
	return k1.(int) < k2.(int)
}

func equal(k1, k2 interface{}) bool {
	return k1 == k2
}

func TestRbtree_Insert(t *testing.T) {
	tree := NewRBTree(less, equal)
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Insert(3, 3)
	tree.Insert(4, 4)
	tree.Insert(5, 5)
	tree.Insert(6, 6)
	tree.Insert(7, 7)
	tree.Insert(8, 8)
	fmt.Println(tree.Search(8))
}
