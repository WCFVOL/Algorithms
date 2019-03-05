package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func less(k1, k2 interface{}) bool {
	return k1.(int) < k2.(int)
}

func equal(k1, k2 interface{}) bool {
	return k1 == k2
}

func TestRbtree_Insert(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tree := NewRBTree(less, equal)
	list := [1000]int{}
	for i := 0; i < 1000; i++ {
		key := rand.Intn(100)
		value := key
		tree.Insert(key, value)
		list[i] = key
	}
	for i := 999; i >= 0; i-- {
		key := list[i]
		value := tree.Search(key)
		if key != value {
			fmt.Println(key, value)
		}
	}
}

func TestRbtree_Remove(t *testing.T) {
	//rand.Seed(time.Now().Unix())
	tree := NewRBTree(less, equal)
	for i := 0; i < 5; i++ {
		key := i
		value := key
		tree.Insert(key, value)
	}
	for i := 0; i < 5; i++ {
		key := i
		fmt.Println(tree.Remove(key))
	}

}
