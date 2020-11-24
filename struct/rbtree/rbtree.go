package rbtree

import (
	"fmt"
	"math"
)

// 红黑树规则
// 1. 节点是红色或黑色。
// 2. 根是黑色。
// 3. 所有叶子都是黑色（叶子是NIL节点）。
// 4. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
// 5. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
// https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91

// RBTree 红黑树
type RBTree struct {
	root *RBLeaf
}

// RBLeaf 红黑节点
type RBLeaf struct {
	key                 int
	left, right, parent *RBLeaf
	red                 bool
}

func (l *RBLeaf) insert(k *RBLeaf) {
	// 绑个父元素
	k.parent = l
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

// 返回叔父节点
func (l *RBLeaf) uncle() *RBLeaf {

	// 无祖父 或 无父 说明是根 或 根的左右子叶
	if l.parent == nil || l.grandpa() == nil {
		return nil
	}

	// 返回叔父
	if l.parent == l.grandpa().left {
		return l.grandpa().right
	}

	return l.grandpa().left

}

// 返回祖父节点
func (l *RBLeaf) grandpa() *RBLeaf {
	// 节点不是根
	if l.parent != nil {
		return l.parent.parent
	}

	return nil
}

// 用目标替换当前节点
func (t *RBTree) replace(old, new *RBLeaf) {
	if old == t.root {
		// 如果要替换的是root
		t.root = new
	} else {
		p := old.parent

		if p.left == old {
			p.left = new
		}

		if p.right == old {
			p.right = new
		}
	}

}

// 左旋
func (t *RBTree) leftRotate(pivot *RBLeaf) *RBTree {
	if pivot == t.root {
		return t
	}
	p := pivot.parent
	// 父节点的右节点 指向目标(pviot)节点的左节点
	p.right = pivot.left
	// 目标(pviot)节点的左节点 指向 父节点
	pivot.left = p
	// 把目标节点替换成父节点
	t.replace(p, pivot)

	return t
}

// 右旋
func (t *RBTree) rightRotate(pivot *RBLeaf) *RBTree {
	if pivot == t.root {
		return t
	}
	p := pivot.parent
	// 父节点的左节点 指向目标(pviot)节点的右节点
	p.left = pivot.right
	// 目标(pviot)节点的右节点指向父节点
	pivot.right = p
	// 把目标节点替换成父节点
	t.replace(p, pivot)
	return t
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

// Find 查找
func (t *RBTree) Find(k int) *RBLeaf {
	cur := t.root
	for cur != nil {
		if k > cur.key {
			cur = cur.right
		} else if k < cur.key {
			cur = cur.left
		} else {
			return cur
		}
	}
	return nil
}

// Insert 插入 DFS
func (t *RBTree) Insert(k int) *RBTree {
	n := &RBLeaf{key: k, red: true}
	if t.root == nil {
		t.root = n
	} else {
		t.root.insert(n)
		// 插入自平衡
	}
	t.insertCheck1(n)

	return t
}

func (t *RBTree) insertCheck1(l *RBLeaf) {
	if l.parent == nil {
		// 如果是根节点
		// 根据: 2. 根是黑色。
		// 设为黑色
		l.red = false
	} else {
		// 不是根节点, 走规则2
		t.insertCheck2(l)
	}
}

func (t *RBTree) insertCheck2(l *RBLeaf) {
	// 因为插入节点是红色
	if l.parent.red == false {
		// 父节点是黑色, ok
		return
	}
	// 不是黑色, 则与: 4. 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
	// 冲突, 走规则3
	t.insertCheck3(l)
}

func (t *RBTree) insertCheck3(l *RBLeaf) {
	// 如果 父(insertCheck2己证明)、叔父都是红色
	if l.uncle() != nil && l.uncle().red == true {
		// 把父、叔、爷节点变色
		l.parent.red = false
		l.uncle().red = false
		l.grandpa().red = true
		// 递归检查 爷节点
		t.insertCheck1(l.grandpa())
	} else {
		t.insertCheck4(l)
	}

}

func (t *RBTree) insertCheck4(l *RBLeaf) {
	// 这个场景是叔节点为空或是黑色
	// 我们先要做一次旋转, 旋转规则如下:
	if l == l.parent.right && l.parent == l.grandpa().left {
		// 如果目标在父节点右边, 且父节点在爷节点左边, 左旋
		//         黑             红
		//        / \     染色    / \
		//       红  黑    ->    黑  黑
		//      /  \           /  \
		//     Nil (红)       Nil (红)
		t.leftRotate(l)
		// 因为左旋, 这里l.left是之前的父节点
		l = l.left
	} else if l == l.parent.left && l.parent == l.grandpa().right {
		// 如果目标在父节点左边, 且父节点在爷节点右边, 右旋
		t.rightRotate(l)
		// 因为右旋, 这里l.right是之前的父节点
		l = l.right
	}

	// 第二次旋转
	t.insertCheck5(l)
}

func (t *RBTree) insertCheck5(l *RBLeaf) {
	// 第二次旋转, 先把父元素染黑、爷元素染红
	l.parent.red = false
	l.grandpa().red = true
	if l == l.parent.left && l.parent == l.grandpa().left {
		// 当目标在父节点左侧且父也在爷左时
		//         黑             红               黑
		//        / \     染色    / \      右旋    / \
		//       红  黑    ->    黑  黑     -->  (红) 红
		//      /               /                     \
		//     (红)           (红)                     黑
		// 右旋其父
		t.rightRotate(l.parent)
	} else {
		// 否则就是述情况的镜像
		// 左旋其父
		t.leftRotate(l.parent)
	}
}

// Remove 删除
// 删除逻辑
// 如果同时有左右子叶, 则用右子叶的最左子叶代替位置(逻辑上 左子叶的最右子叶也可以),
// 否则用存在的子叶代替
func (t *RBTree) Remove(k int) *RBTree {
	l := t.Find(k)
	if l == nil {
		return t
	}
	if l.left == nil && l.right == nil {
		// 左右子叶都为空, 用nil代替
		t.replace(l, nil)
	} else if l.left != nil && l.right != nil {
		// 左右子叶都不为空
		// 找到右子叶的最左子叶节点
		cur := l.right
		for cur.left != nil {
			cur = cur.left
		}

		// 从父节点用它的右树替换它
		t.replace(cur, cur.right)
		// 把原节点的左右子叶赋给目标
		cur.left = l.left
		cur.right = l.right
		// 替换他
		t.replace(l, cur)

	} else if l.left == nil {
		// 左子叶为空
		t.replace(l, l.right)
	} else {
		// 右子叶为空
		t.replace(l, l.left)
	}

	return t
}
