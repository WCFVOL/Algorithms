package rbtree

type rbtree struct {
	root  *node
	size  int
	less  func(k1, k2 interface{}) bool
	equal func(k1, k2 interface{}) bool
}

func NewRBTree(less func(a, b interface{}) bool, equal func(a, b interface{}) bool) *rbtree {
	result := &rbtree{
		less:  less,
		equal: equal,
	}
	return result
}

func (tree *rbtree) Insert(k, v interface{}) {
	nod := newNode(k, v)
	nod.left = nil
	nod.right = nil
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
	now := tree.searchNode(k)
	if now == nil {
		return nil
	} else {
		return now.v
	}
}

func (tree *rbtree) searchNode(k interface{}) *node {
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
		return now
	}
}

func (tree *rbtree) Remove(k interface{}) interface{} {
	now := tree.searchNode(k)
	if now == nil {
		return nil
	}
	oldValue := now.v
	tree.removeNode(now)
	return oldValue
}

func (tree *rbtree) removeNode(p *node) {
	if p.left != nil && p.right != nil {
		s := successor(p)
		p.k = s.k
		p.v = s.v
		p = s
	}
	var rep *node
	if p.left != nil {
		rep = p.left
	} else {
		rep = p.right
	}
	if rep != nil {
		rep.p = p.p
		if p.p == nil {
			tree.root = rep
		} else if p == p.p.left {
			p.p.left = rep
		} else {
			p.p.right = rep
		}
		p.left = nil
		p.right = nil
		p.p = nil
		if !p.red {
			tree.deleteFixUp(rep)
		}
	} else if p.p == nil {
		tree.root = nil
	} else {
		if !p.red {
			tree.deleteFixUp(p)
		}
		if p.p != nil {
			if p == p.p.left {
				p.p.left = nil
			} else if p == p.p.right {
				p.p.right = nil
			}
			p.p = nil
		}
	}
}

func (tree *rbtree) deleteFixUp(now *node) {
	for now != tree.root && !now.Red() {
		if now == now.Parent().Left() {
			w := now.Parent().Right()
			if w.Red() {
				w.SetRed(false)
				now.Parent().SetRed(true)
				now.Parent().leftRotate(tree)
				w = now.Parent().Right()
			}
			if !w.Right().Red() && !w.Left().Red() {
				w.SetRed(true)
				now = now.Parent()
			} else {
				if !w.Right().Red() {
					w.SetRed(true)
					w.Left().SetRed(false)
					w.rightRotate(tree)
					w = now.Parent().Right()
				}
				w.SetRed(now.Parent().Red())
				now.Parent().SetRed(false)
				w.Right().SetRed(false)
				now.Parent().leftRotate(tree)
				now = tree.root
			}
		} else {
			w := now.Parent().Left()
			if w.Red() {
				w.SetRed(false)
				now.Parent().SetRed(true)
				now.Parent().rightRotate(tree)
				w = now.Parent().Left()
			}
			if !w.Left().Red() && !w.Right().Red() {
				w.SetRed(true)
				now = now.Parent()
			} else {
				if !w.Left().Red() {
					w.SetRed(true)
					w.Right().SetRed(false)
					w.leftRotate(tree)
					w = now.Parent().Left()
				}
				w.SetRed(now.Parent().Red())
				now.Parent().SetRed(false)
				w.Left().SetRed(false)
				now.Parent().rightRotate(tree)
				now = tree.root
			}
		}
	}
	now.red = false
}

// 把u父亲的儿子u 替换为v 不涉及u和v的孩子
func (tree *rbtree) transplant(u *node, v *node) {
	if u.p == nil {
		tree.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (tree *rbtree) Clear() {
	tree.root = nil
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
				nod.p.p.red = true
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
				nod.p.p.red = true
				nod.p.p.leftRotate(tree)
			}
		}
	}
	tree.root.red = false
}

func predecessor(t *node) *node {
	if t == nil {
		return nil
	} else if t.left != nil {
		p := t.left
		for p.right != nil {
			p = p.right
		}
		return p
	} else {
		p := t.p
		ch := t
		for p != nil && ch == p.left {
			ch = p
			p = p.p
		}
		return p
	}
}

func successor(t *node) *node {
	if t == nil {
		return nil
	} else if t.right != nil {
		p := t.right
		for p.left != nil {
			p = p.left
		}
		return p
	} else {
		p := t.p
		ch := t
		for p != nil && ch == p.right {
			ch = p
			p = p.p
		}
		return p
	}
}
