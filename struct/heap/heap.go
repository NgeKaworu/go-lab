package heap

// Heap 堆
type Heap struct {
	heap []int
}

// Heapify 堆化
func Heapify(arr []int) *Heap {
	h := new(Heap)
	for _, v := range arr {
		h.Insert(v)
	}
	return h
}

// Len 长度
func (h *Heap) Len() int {
	return h.Len()
}

// Insert 插入
func (h *Heap) Insert(val int) *Heap {
	h.heap = append(h.heap, val)
	h.swim(h.Len())
	return h
}

// Pop 弹出
func (h *Heap) Pop() (res int) {
	res = h.heap[0]
	h.heap[0] = h.heap[h.Len()-1]
	h.heap = h.heap[:h.Len()-1]
	h.sink(0)
	return
}

// parent 返回父节点
func (h *Heap) parent(pos int) int {
	return (pos - 1) / 2
}

// left 返回左节点
func (h *Heap) left(pos int) int {
	return pos*2 + 1

}

// right 返回左节点
func (h *Heap) right(pos int) int {
	return pos*2 + 2

}

// swap 交换位置
func (h *Heap) swap(p1, p2 int) *Heap {
	h.heap[p1], h.heap[p2] = h.heap[p2], h.heap[p1]
	return h
}

// Swim 上浮
func (h *Heap) swim(pos int) *Heap {
	for pos > 0 && h.heap[pos] > h.heap[h.parent(pos)] {
		// 如果第 k 个元素不是堆顶且比上层大
		// 将 k 换上去
		h.swap(pos, h.parent(pos))
		pos = h.parent(pos)
	}
	return h
}

// Sink 下沉
func (h *Heap) sink(pos int) *Heap {
	l := h.Len()
	// 如果沉到堆底，就沉不下去了
	for h.left(pos) < l {
		// 先假设左边节点较大
		t := h.left(pos)
		// 如果右边节点存在，比一下大小
		if h.right(pos) < l && h.heap[t] < h.heap[h.right(pos)] {
			t = h.right(pos)
		}
		// 结点 k 比俩孩子都大，就不必下沉了
		if h.heap[t] < h.heap[pos] {
			break
		}
		// 否则，不符合最大堆的结构，下沉 k 结点
		h.swap(pos, t)
		pos = t
	}

	return h
}
