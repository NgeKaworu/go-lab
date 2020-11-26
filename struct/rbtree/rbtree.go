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
// PS: 本文中我们使用"nil叶子"或"空（null）叶子".
// https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91

// RBTree 红黑树
type RBTree struct {
	root *TreeNode
}

// TreeNode 树节点
type TreeNode struct {
	key                 int
	left, right, parent *TreeNode
	red                 bool
}

func (n *TreeNode) insert(k *TreeNode) {
	// 绑个父元素
	k.parent = n
	// DFS 深度优先
	if k.key <= n.key {
		// 小于等于走左边
		if n.left == nil {
			// 左树为空, 左树即是节点
			n.left = k
		} else {
			// 递归左树
			n.left.insert(k)
		}
	} else {
		// 大于走右边
		if n.right == nil {
			// 右树为空, 右树即节点
			n.right = k
		} else {
			// 递归右树
			n.right.insert(k)
		}
	}
}

// 返回叔父节点
func (n *TreeNode) uncle() *TreeNode {

	// 无祖父 或 无父 说明是根 或 根的左右子树
	if n.parent == nil || n.grandpa() == nil {
		return nil
	}

	// 返回叔父
	if n.parent == n.grandpa().left {
		return n.grandpa().right
	}

	return n.grandpa().left

}

// 返回兄弟节点
func (n *TreeNode) sibling() *TreeNode {
	p := n.parent
	if p != nil {
		if p.left == n {
			return p.right
		}
		return p.left
	}
	return nil
}

// 返回祖父节点
func (n *TreeNode) grandpa() *TreeNode {
	// 节点不是根
	if n.parent != nil {
		return n.parent.parent
	}

	return nil
}

// 用目标替换当前节点
func (t *RBTree) replace(old, new *TreeNode) {
	if new != nil {
		// 双向链 记得 改绑 父元素
		new.parent = old.parent
	}

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
func (t *RBTree) leftRotate(pivot *TreeNode) *RBTree {
	if pivot == t.root {
		return t
	}
	p := pivot.parent
	// 父节点的右节点 指向目标(pviot)节点的左节点
	p.right = pivot.left
	// 记得双向绑定
	if pivot.left != nil {
		pivot.left.parent = p
	}
	// 目标(pviot)节点的左节点 指向 父节点
	pivot.left = p
	// 处理双向绑定
	pivot.parent = nil
	// 用目标节点替换成父节点
	t.replace(p, pivot)
	// 最后还需要在处理一次双向绑定
	p.parent = pivot
	return t
}

// 右旋
func (t *RBTree) rightRotate(pivot *TreeNode) *RBTree {
	if pivot == t.root {
		return t
	}
	p := pivot.parent
	// 父节点的左节点 指向目标(pviot)节点的右节点
	p.left = pivot.right
	if pivot.right != nil {
		// 处理双向绑定
		pivot.right.parent = p
	}
	// 目标(pviot)节点的右节点指向父节点
	pivot.right = p
	// 处理双向绑定
	pivot.parent = nil
	// 用目标节点替换成父节点
	t.replace(p, pivot)
	// 最后还需要在处理一次双向绑定
	p.parent = pivot
	return t
}

// BFS 广度优先 打印
func (t *RBTree) String() (str string) {
	// 初始化队列
	queue := []*TreeNode{t.root}
	count, line := 0.0, 1.0
	for true {
		// 退出条件
		// 如果某一行的下一行都是空
		if float64(len(queue)) == math.Pow(2, line) {
			every := true
			for _, v := range queue {
				if v != nil {
					every = false
					break
				}
			}
			if every == true {
				break
			}
		}
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
			// 保证格式一致
			queue = append(queue, nil, nil)
		}

	}
	return
}

// Find 查找
func (t *RBTree) Find(k int) *TreeNode {
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
	n := &TreeNode{key: k, red: true}
	if t.root == nil {
		t.root = n
	} else {
		t.root.insert(n)
		// 插入自平衡
	}
	t.insertCheck(n)

	return t
}

// 插入检查
func (t *RBTree) insertCheck(n *TreeNode) {
	if n.parent == nil {
		// 如果是根节点
		// 根据: 2. 根是黑色。
		// 设为黑色
		n.red = false
		return

	}

	// 不是根节点的逻辑
	// 因为插入节点是红色
	// 父节点是红色
	// 则与规则4冲突 —— 每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
	if n.parent.red == true {
		// 如果 父(insertCheck2己证明)、叔父都是红色
		if n.uncle() != nil && n.uncle().red == true {
			// 把父、叔、爷节点变色
			n.parent.red = false
			n.uncle().red = false
			n.grandpa().red = true
			// 递归检查 爷节点
			t.insertCheck(n.grandpa())
		} else {
			// 这个场景是叔节点为空或是黑色
			// 假设我们的树不是“顺拐” 要先做一次旋转, 旋转规则如下:
			if n == n.parent.right && n.parent == n.grandpa().left {
				// 如果目标在父节点右边, 且父节点在爷节点左边, 左旋
				//         黑             红
				//        / \     染色    / \
				//       红  黑    ->    黑  黑
				//      /  \           /  \
				//     Nil (红)       Nil (红)
				t.leftRotate(n)
				// 因为左旋, 这里n.left是之前的父节点
				n = n.left
			} else if n == n.parent.left && n.parent == n.grandpa().right {
				// 如果目标在父节点左边, 且父节点在爷节点右边, 右旋
				t.rightRotate(n)
				// 因为右旋, 这里n.right是之前的父节点
				n = n.right
			}

			// 第二次旋转, 先把父元素染黑、爷元素染红
			n.parent.red = false
			n.grandpa().red = true

			if n == n.parent.left && n.parent == n.grandpa().left {
				// 当目标在父节点左侧且父也在爷左时
				//         黑             红               黑
				//        / \     染色    / \      右旋    / \
				//       红  黑    ->    黑  黑     -->  (红) 红
				//      /               /                     \
				//     (红)           (红)                     黑
				// 右旋其父
				t.rightRotate(n.parent)
			} else {
				// 否则就是述情况的镜像
				// 左旋其父
				t.leftRotate(n.parent)
			}
		}
	}

}

// Remove 删除
// 删除逻辑
// 如果同时有左右子树, 则用右的最左子树代替位置(逻辑上 左子树的最右子树也可以),
// 否则用存在的子树代替
func (t *RBTree) Remove(k int) *RBTree {
	n := t.Find(k)
	if n == nil {
		return t
	}
	if n.left == nil && n.right == nil {
		// 左右子树都为空, 用nil代替
		t.replace(n, nil)
	} else if n.left != nil && n.right != nil {
		// 左右子树都不为空
		// 找到右子树的最左子树节点
		cur := n.right
		for cur.left != nil {
			cur = cur.left
		}

		// 从父节点用它的右树替换它
		t.replace(cur, cur.right)
		// 把原节点的左右子树赋给目标
		cur.left = n.left
		cur.right = n.right
		// 替换他
		t.replace(n, cur)

	} else if n.left == nil {
		// 左子树为空
		t.replace(n, n.right)
	} else {
		// 右子树为空
		t.replace(n, n.left)
	}

	return t
}
