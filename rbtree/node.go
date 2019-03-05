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

func newNil() *node {
	return &node{
		left:  nil,
		right: nil,
		red:   false,
		k:     nil,
		v:     nil,
	}
}

func (now *node) leftRotate(tree *rbtree) {
	right := now.right
	now.right = right.left
	if right.left != tree.nil {
		right.left.p = now
	}
	right.p = now.p
	if now.p != tree.nil {
		if now.p.left == now {
			now.p.left = right
		} else {
			now.p.right = right
		}
	} else {
		tree.root = right
	}
	right.left = now
	now.p = right
}

func (now *node) rightRotate(tree *rbtree) {
	left := now.left
	now.left = left.right
	if left.right != tree.nil {
		left.right.p = now
	}
	left.p = now.p
	if now.p != tree.nil {
		if now.p.left == now {
			now.p.left = left
		} else {
			now.p.right = left
		}
	} else {
		tree.root = left
	}
	left.right = now
	now.p = left
}
