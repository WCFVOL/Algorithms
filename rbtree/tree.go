package rbtree

type Tree interface {
	Insert() interface{}
	Search(k interface{}) interface{}
	Remove(k interface{}) interface{}
	Clear()
}
