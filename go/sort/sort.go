package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func siftDown(e Interface, l, r, first int) {
	root := l
	for {
		child := 2*root + 1
		if child >= r {
			return
		}
		if child+1 < r && e.Less(first+child, first+child+1) {
			child++
		}
		if !e.Less(root, child) {
			return
		}
		e.Swap(root, child)
		root = child
	}
}

// heap sort for (l,r)
func heapSort(e Interface, l, r int) {
	first := l
	lo := 0
	hi := r - l
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(e, i, hi, first)
	}
	// pop element
	for i := hi - 1; i >= 0; i-- {
		e.Swap(first, first+i)
		siftDown(e, lo, i, first)
	}
}
