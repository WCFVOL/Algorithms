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

func (now *node) leftRotate() {
	right := now.right
	now.right = right.left
	if right.left != nil {
		right.left.p = now
	}
	right.p = now.p
	if now.p != nil {
		if now.p.left == now {
			now.p.left = right
		} else {
			now.p.right = right
		}
	}
	right.left = now
	now.p = right
}

func (now *node) rightRotate() {
	left := now.left
	now.left = left.right
	if left.right != nil {
		left.right.p = now
	}
	left.p = now.p
	if now.p != nil {
		if now.p.left == now {
			now.p.left = left
		} else {
			now.p.right = left
		}
	}
	left.right = now
	now.p = left
}
