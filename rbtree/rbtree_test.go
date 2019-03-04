package rbtree

import (
	"fmt"
	"math/rand"
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
	list := [100]int{}
	for i := 0; i < 100; i++ {
		key := rand.Intn(100)
		value := key
		tree.Insert(key, value)
		list[i] = key
	}
	for i := 99; i >= 0; i-- {
		fmt.Print(list[i], ",")
		fmt.Println(tree.Search(list[i]))
	}
}
