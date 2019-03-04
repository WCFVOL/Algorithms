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
	for i := 0; i < 100; i++ {

	}
	fmt.Println(tree.Search(6))
}
