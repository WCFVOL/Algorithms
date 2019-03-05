package rbtree

type rbtree struct {
	root  *node
	size  int
	nil   *node
	less  func(k1, k2 interface{}) bool
	equal func(k1, k2 interface{}) bool
}

func NewRBTree(less func(a, b interface{}) bool, equal func(a, b interface{}) bool) *rbtree {
	result := &rbtree{
		less:  less,
		equal: equal,
		nil:   newNil(),
	}
	result.root = result.nil
	return result
}

func (tree *rbtree) Insert(k, v interface{}) {
	nod := newNode(k, v)
	nod.left = tree.nil
	nod.right = tree.nil
	now := tree.root
	fa := tree.nil
	for now != tree.nil {
		fa = now
		if tree.less(nod.k, fa.k) {
			now = fa.left
		} else {
			now = fa.right
		}
	}
	nod.p = fa
	if fa == tree.nil {
		tree.root = nod
	} else if tree.less(nod.k, fa.k) {
		fa.left = nod
	} else {
		fa.right = nod
	}
	tree.insertFixUp(nod)
}

func (tree *rbtree) Search(k interface{}) interface{} {
	now := tree.searchNode(k)
	if now == tree.nil {
		return nil
	} else {
		return now.v
	}
}

func (tree *rbtree) searchNode(k interface{}) *node {
	now := tree.root
	if now == tree.nil {
		return nil
	}
	for now != tree.nil && !tree.equal(k, now.k) {
		if tree.less(k, now.k) {
			now = now.left
		} else {
			now = now.right
		}
	}
	if now == tree.nil {
		return nil
	} else {
		return now
	}
}

func (tree *rbtree) Remove(k interface{}) interface{} {
	now := tree.searchNode(k)
	if now == tree.nil {
		return nil
	}
	tree.removeNode(now)
	return now.v
}

func (tree *rbtree) removeNode(nod *node) {
	now := nod
	var x *node
	red := now.red
	if now.left == tree.nil {
		x = now.right
		tree.transplant(now, x)
	} else if now.right == tree.nil {
		x = now.left
		tree.transplant(now, x)
	} else {
		now = tree.Minimum(now.right)
		red = now.red
		x = now.right
		if now.p == nod {
			x.p = now
		} else {
			tree.transplant(now, now.right)
			now.right = nod.right
			now.right.p = now
		}
		tree.transplant(nod, now)
		now.left = nod.left
		now.left.p = now
		now.red = nod.red
	}
	if !red {
		tree.deleteFixUp(x)
	}
}

func (tree *rbtree) deleteFixUp(nod *node) {
	//defer func(tree *node) {
	//	fmt.Println("wtf")
	//}(nod)
	now := nod
	for now != tree.root && !now.red {
		if now == now.p.left {
			w := now.p.right
			if w.red {
				w.red = false
				now.p.red = true
				now.p.leftRotate(tree)
				w = now.p.right
			}
			if !w.right.red && !w.left.red {
				w.red = true
				now = now.p
			} else if w.left.red {
				w.red = true
				w.left.red = false
				w.rightRotate(tree)
			} else {
				w.red = now.p.red
				now.p.red = false
				w.right.red = false
				now.p.leftRotate(tree)
				now = tree.root
			}
		} else {
			w := now.p.left
			if w.red {
				w.red = false
				now.p.red = true
				now.p.rightRotate(tree)
				w = now.p.left
			}
			if !w.right.red && !w.left.red {
				w.red = true
				now = now.p
			} else if w.right.red {
				w.red = true
				w.right.red = false
				w.leftRotate(tree)
			} else {
				w.red = now.p.red
				now.p.red = false
				w.left.red = false
				now.p.rightRotate(tree)
				now = tree.root
			}
		}
	}
	now.red = false
}

// 把u父亲的儿子u 替换为v 不涉及u和v的孩子
func (tree *rbtree) transplant(u *node, v *node) {
	if u.p == tree.nil {
		tree.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (tree *rbtree) Clear() {
	tree.root = tree.nil
}

func (tree *rbtree) insertFixUp(now *node) {
	nod := now
	for nod.p != tree.nil && nod.p.red {
		if nod.p == nod.p.p.left {
			y := nod.p.p.right
			if y != tree.nil && y.red {
				nod.p.red = false
				y.red = false
				nod.p.p.red = true
			} else if nod == nod.p.right {
				nod = nod.p
				nod.leftRotate(tree)
			} else {
				nod.p.red = false
				nod.p.p.red = true
				nod.p.p.rightRotate(tree)
			}
		} else {
			y := nod.p.p.left
			if y != tree.nil && y.red {
				nod.p.red = false
				y.red = false
				nod.p.p.red = true
			} else if nod == nod.p.left {
				nod = nod.p
				nod.rightRotate(tree)
			} else {
				nod.p.red = false
				nod.p.p.red = true
				nod.p.p.leftRotate(tree)
			}
		}
	}
	tree.root.red = false
}

//nod子树的最小值
func (tree *rbtree) Minimum(nod *node) *node {
	now := nod
	if now == tree.nil {
		return nil
	}
	for now.left != tree.nil {
		now = now.left
	}
	return now
}

//nod子树的最大值
func (tree *rbtree) Maximum(nod *node) *node {
	now := nod
	if now == tree.nil {
		return nil
	}
	for now.left != tree.nil {
		now = now.left
	}
	return now
}
