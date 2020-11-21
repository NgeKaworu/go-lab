package chain

import (
	"log"
)

// Chain 链表
type Chain struct {
	Head *ListNode
}

// Len 返回长度
func (c *Chain) Len() (l int) {
	cursor := c.Head
	for cursor != nil {
		l++
		cursor = cursor.Next
	}
	return
}

// FromArr 数组转链表
func FromArr(arr []int) *Chain {
	head := new(ListNode)
	cur := head
	for _, v := range arr {
		cur.Val = v
		cur.Next = new(ListNode)
		cur = cur.Next
	}
	return &Chain{head}
}

// ToArr 链表转数组
func (c *Chain) ToArr() (arr []int) {
	cursor := c.Head
	for cursor.Next != nil {
		arr = append(arr, cursor.Val)
		cursor = cursor.Next
	}
	return
}

func (c *Chain) String() (s string) {
	cur := c.Head
	for cur != nil {
		s += "val -> "
		cur = cur.Next
	}
	return
}

// InsertH 头插法
func (c *Chain) InsertH(val int) *Chain {
	l := &ListNode{val, c.Head}
	c.Head = l
	return c
}

// InsertT 尾插法
func (c *Chain) InsertT(val int) *Chain {
	l := &ListNode{val, nil}
	cursor := c.Head
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = l
	return c
}

// Insert 增
func (c *Chain) Insert(pos, val int) *Chain {
	l := c.Find(pos)
	if l != nil {
		l.Insert(val)
	} else {
		log.Fatalf("not find pos as %v", pos)
	}
	return c
}

// Remove 删除
func (c *Chain) Remove(pos int) *Chain {
	pre := c.Head
	for pre != nil && pre.Next != nil && pre.Next.Val != pos {
		pre = pre.Next
	}
	if pre != nil {
		// pre.Next是目标节点, pre.Next.Next是目标节点对下一节点的引用,设置为空可以有效出发垃圾回收
		pre.Next, pre.Next.Next = pre.Next.Next, nil
	} else {
		log.Fatalf("not find pos as %v", pos)
	}
	return c
}

// Update 改
func (c *Chain) Update(pos, val int) *Chain {
	l := c.Find(pos)
	if l != nil {
		l.Val = val
	} else {
		log.Fatalf("not find pos as %v", pos)
	}
	return c
}

// Find 查
func (c *Chain) Find(val int) *ListNode {
	cur := c.Head
	for cur != nil && cur.Val != val {
		cur = cur.Next
	}

	if cur == nil {
		log.Fatalf("not find node as %v", val)
	}

	return cur
}

// Swap 交换节点位置
func (c *Chain) Swap(n1, n2 int) *Chain {
	if n1 == n2 {
		log.Fatalf("same node")
	}

	cur := c.Head

	if n1 == cur.Val || n2 == cur.Val {
		log.Fatalf("Head node swap please use HeadSwap method")
	}

	var l1, l2 *ListNode
	// 找到要交换节点的前一个节点
	for cur.Next != nil && (l1 == nil || l2 == nil) {
		if cur.Next.Val == n1 {
			l1 = cur
		}
		if cur.Next.Val == n2 {
			l2 = cur
		}
	}

	if l1 == nil || l2 == nil {
		log.Fatalf("n1 or n2 not found")
	}

	// 下面建议画图理解
	t := l1.Next.Next
	l1.Next.Next = l2.Next.Next
	l2.Next.Next = t
	t = l2.Next
	l2.Next = l1.Next
	l1.Next = t

	return c
}

// HeadSwap 交换头节点和目标节点位置
func (c *Chain) HeadSwap(val int) *Chain {

	cur := c.Head
	if val == cur.Val {
		log.Fatalf("%v and Head is same", val)
	}

	var l *ListNode
	// 找到要交换节点的前一个节点
	for cur.Next != nil && l == nil {
		if cur.Next.Val == val {
			l = cur
		}
	}

	if l == nil {
		log.Fatal("val not found")
	}

	t := l.Next
	c.Head.Next, l.Next, t.Next = nil, c.Head, c.Head.Next

	return &Chain{t}
}

// Sort 并归排序
func (c *Chain) Sort() *Chain {
	head := c.Head
	// 0 or 1 element.
	if head == nil || head.Next == nil {
		return c
	}
	// 2 pointers, if the fast point comes to the end, the slow one must be in the middle.
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// split into 2 parts.
	n := slow.Next
	slow.Next = nil
	// Sort recursivley
	merge := func(c1 *Chain, c2 *Chain) *Chain {
		node1, node2 := c1.Head, c2.Head
		// Create a new empty list.
		node := &ListNode{Val: 0}
		current := node
		// Compare one by one, put the smaller value into the new list.
		for node1 != nil && node2 != nil {
			if node1.Val <= node2.Val {
				current.Next = node1
				node1 = node1.Next
			} else {
				current.Next = node2
				node2 = node2.Next
			}
			current = current.Next
		}
		if node1 != nil {
			current.Next = node1
			node1 = node1.Next
		}
		if node2 != nil {
			current.Next = node2
			node2 = node2.Next
		}
		return &Chain{node.Next}
	}
	return merge((&Chain{head}).Sort(), (&Chain{n}).Sort())
}

// ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// Insert 向节点后插入
func (l *ListNode) Insert(val int) *ListNode {
	n := &ListNode{val, l.Next}
	l.Next = n
	return l
}
