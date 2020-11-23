package rbtree

import (
	"fmt"
	"math"
)

// RBTree 红黑树
type RBTree struct {
	root *RBLeaf
}

// RBLeaf 红黑节点
type RBLeaf struct {
	key         int
	left, right *RBLeaf
	red         bool
}

func (l *RBLeaf) insert(k *RBLeaf) {
	// DFS 深度优先
	if k.key <= l.key {
		// 小于等于走左边
		if l.left == nil {
			// 左树为空, 左树即是节点
			l.left = k
		} else {
			// 递归左树
			l.left.insert(k)
		}
	} else {
		// 大于走右边
		if l.right == nil {
			// 右树为空, 右树即节点
			l.right = k
		} else {
			// 递归右树
			l.right.insert(k)
		}
	}
}

// BFS 广度优先 打印
func (t *RBTree) String() (str string) {
	// 初始化队列
	queue := []*RBLeaf{t.root}
	count, line := 0.0, 1.0
	for len(queue) != 0 {
		// 插入换行
		// 高度h = h(n) = log2(n) + 1, n是节点数
		h := math.Floor(math.Log2(count + 1))
		if h == line {
			str += "\n"
			line++
		}
		count++

		// pop
		leaf := queue[0]
		queue = queue[1:]
		if leaf != nil {
			str += fmt.Sprintf("(%v, %v) ", leaf.key, leaf.red)
			queue = append(queue, leaf.left, leaf.right)
		} else {
			str += " nil "
		}

	}
	return
}

// Insert 插入 DFS
func (t *RBTree) Insert(k int) *RBTree {
	n := &RBLeaf{key: k, red: true}
	if t.root == nil {
		t.root = n
	} else {
		t.root.insert(n)
	}

	return t
}

// Remove 删除
func (t *RBTree) Remove(k int) *RBTree {

	return t
}
