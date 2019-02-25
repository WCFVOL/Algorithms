package rbtree

type rbtree struct {
	root *node
	size int
	less func(a,b interface{}) bool
}

func NewRBTree(less func(a,b interface{}) bool) *rbtree{
	return &rbtree{
		less:less,
	}
}