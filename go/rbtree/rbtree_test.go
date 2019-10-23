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
	list := [10000]int{}
	v := 0
	for i := 0; i < 10000; i++ {
		key := rand.Intn(10000)
		value := key
		tree.Insert(key, value)
		r := isRBTree(tree)
		if r > v {
			v = r
		}
		list[i] = key
	}
	fmt.Println(v)
	for i := 9999; i >= 0; i-- {
		key := list[i]
		value := tree.Search(key)
		if key != value {
			fmt.Println(key, value)
		}
	}
}

func TestRbtree_Remove(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tree := NewRBTree(less, equal)
	list := [10000]int{}
	for i := 0; i < 10000; i++ {
		key := rand.Intn(10000)
		value := key
		tree.Insert(key, value)
		list[i] = key
		fmt.Print(key, ",")
	}
	fmt.Println()
	v := 0
	for i := 9999; i >= 0; i-- {
		key := list[i]
		value := tree.Remove(key)
		r := isRBTree(tree)
		if r > v {
			v = r
		}
		if key != value {
			fmt.Println(key, value)
		}
	}
	fmt.Println(v)
}

func isRBTree(tree *rbtree) int {
	list := []int{}
	//ok := true
	if tree.root == nil {
		return 0
	} else {
		dfs(tree, tree.root, &list, 0)
		max := 0
		min := 999999
		for i := 1; i < len(list); i++ {
			if list[i] > max {
				max = list[i]
			}
			if list[i] < min {
				min = list[i]
			}
		}
		return max - min
	}

}

func dfs(tree *rbtree, nod *node, list *[]int, cnt int) {
	if nod == nil {
		*list = append(*list, cnt)
		return
	}
	if !nod.Red() {
		cnt = cnt + 1
	}
	dfs(tree, nod.left, list, cnt)
	dfs(tree, nod.right, list, cnt)
}
