package heap

import (
	"github.com/WCFVOL/Algorithms/sort"
)

type Heap interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Heap) {
	// heap.Init()
	len := h.Len()
	for i := len / 2; i >= 0; i-- {
		down(h, i, len)
	}
}

func Push(h Heap, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h Heap) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(h Heap, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(n, i)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

func down(h Heap, i0, n int) bool {
	i := i0
	for {
		l := 2*i + 1
		if l >= n || l < 0 {
			break
		}
		largest := l
		if r := l + 1; r < n && r >= 0 && h.Less(r, l) {
			largest = r
		}
		if h.Less(i, largest) {
			break
		}
		h.Swap(i, largest)
		i = largest
	}
	return i > i0
}

func up(h Heap, i int) {
	for {
		fa := (i - 1) / 2
		if i == fa || !h.Less(i, fa) {
			break
		}
		h.Swap(i, fa)
		i = fa
	}
}
