package rbtree

type node struct {
	p, left, right *node
	red            bool
	k, v           interface{}
}

func newNode(k, v interface{}) *node {
	return &node{
		left:  nil,
		right: nil,
		red:   true,
		k:     k,
		v:     v,
	}
}

func (p *node) leftRotate(tree *rbtree) {
	if p != nil {
		r := p.right
		p.right = r.left
		if r.left != nil {
			r.left.p = p
		}
		r.p = p.p
		if p.p == nil {
			tree.root = r
		} else if p.p.left == p {
			p.p.left = r
		} else {
			p.p.right = r
		}
		r.left = p
		p.p = r
	}
}

func (p *node) rightRotate(tree *rbtree) {
	if p != nil {
		l := p.left
		p.left = l.right
		if l.right != nil {
			l.right.p = p
		}
		l.p = p.p
		if p.p == nil {
			tree.root = l
		} else if p.p.right == p {
			p.p.right = l
		} else {
			p.p.left = l
		}
		l.right = p
		p.p = l
	}
}

func (nod *node) Red() bool {
	if nod == nil {
		return false
	} else {
		return nod.red
	}
}

func (nod *node) SetRed(red bool) {
	if nod != nil {
		nod.red = red
	}
}

func (nod *node) Right() *node {
	if nod == nil {
		return nil
	} else {
		return nod.right
	}
}

func (nod *node) SetRight(right *node) {
	if nod != nil {
		nod.right = right
	}
}

func (nod *node) Left() *node {
	if nod == nil {
		return nil
	} else {
		return nod.left
	}
}

func (nod *node) SetLeft(left *node) {
	if nod != nil {
		nod.left = left
	}
}

func (nod *node) Parent() *node {
	if nod == nil {
		return nil
	} else {
		return nod.p
	}
}

func (nod *node) SetParent(p *node) {
	if nod != nil {
		nod.p = p
	}
}
