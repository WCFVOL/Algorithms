package rbtree

type node struct {
	p,left,right *node
	red bool
	k,v interface{}
}

func newNode(k,v interface{}) *node{
	return &node{
		red: true,
		k:k,
		v:v,
	}
}

func (now *node) leftRotate() {

}

func (now *node) rightRotate() {

}