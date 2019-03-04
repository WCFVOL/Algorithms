package rbtree

type rbtree struct {
	root  *node
	size  int
	less  func(k1, k2 interface{}) bool
	equal func(k1, k2 interface{}) bool
}

func NewRBTree(less func(a, b interface{}) bool, equal func(a, b interface{}) bool) *rbtree {
	return &rbtree{
		less:  less,
		equal: equal,
	}
}

func (tree *rbtree) Insert(k, v interface{}) {
	nod := newNode(k, v)
	now := tree.root
	var fa *node
	for now != nil {
		fa = now
		if tree.less(nod.k, fa.k) {
			now = fa.left
		} else {
			now = fa.right
		}
	}
	nod.p = fa
	if fa == nil {
		tree.root = nod
	} else if tree.less(nod.k, fa.k) {
		fa.left = nod
	} else {
		fa.right = nod
	}
	tree.insertFixUp(nod)
}

func (tree *rbtree) Search(k interface{}) interface{} {
	now := tree.root
	if now == nil {
		return nil
	}
	for now != nil && !tree.equal(k, now.k) {
		if tree.less(k, now.k) {
			now = now.left
		} else {
			now = now.right
		}
	}
	if now == nil {
		return nil
	} else {
		return now.v
	}
}

func (tree *rbtree) insertFixUp(now *node) {
	nod := now
	for nod.p != nil && nod.p.red {
		if nod.p == nod.p.p.left {
			y := nod.p.p.right
			if y != nil && y.red {
				nod.p.red = false
				y.red = false
				nod.p.p.red = true
			} else if nod == nod.p.right {
				nod = nod.p
				nod.leftRotate(tree)
			} else {
				nod.p.red = false
				nod.p.p.red = false
				nod.p.p.rightRotate(tree)
			}
		} else {
			y := nod.p.p.left
			if y != nil && y.red {
				nod.p.red = false
				y.red = false
				nod.p.p.red = true
			} else if nod == nod.p.left {
				nod = nod.p
				nod.rightRotate(tree)
			} else {
				nod.p.red = false
				nod.p.p.red = false
				nod.p.p.leftRotate(tree)
			}
		}
	}
	tree.root.red = false
}
