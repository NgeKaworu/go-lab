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
func FromArr(arr []interface{}) *Chain {
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
func (c *Chain) ToArr() (arr []interface{}) {
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
func (c *Chain) InsertH(val interface{}) *Chain {
	l := &ListNode{val, c.Head}
	c.Head = l
	return c
}

// InsertT 尾插法
func (c *Chain) InsertT(val interface{}) *Chain {
	l := &ListNode{val, nil}
	cursor := c.Head
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = l
	return c
}

// Insert 增
func (c *Chain) Insert(pos, val interface{}) *Chain {
	l := c.Find(pos)
	if l != nil {
		l.Insert(&ListNode{val, l.Next})
	} else {
		log.Fatalf("not find pos as %v", pos)
	}
	return c
}

// Remove 删除
func (c *Chain) Remove(pos interface{}) *Chain {
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
func (c *Chain) Update(pos, val interface{}) *Chain {
	l := c.Find(pos)
	if l != nil {
		l.Val = val
	} else {
		log.Fatalf("not find pos as %v", pos)
	}
	return c
}

// Find 查
func (c *Chain) Find(val interface{}) *ListNode {
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
func (c *Chain) Swap(n1, n2 interface{}) *Chain {
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
func (c *Chain) HeadSwap(val interface{}) *Chain {

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
func Sort(head *ListNode) *Chain {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}

	merge := func(head1, head2 *ListNode) *ListNode {
		dummyHead := &ListNode{}
		temp, temp1, temp2 := dummyHead, head1, head2
		for temp1 != nil && temp2 != nil {
			if temp1.Val <= temp2.Val {
				temp.Next = temp1
				temp1 = temp1.Next
			} else {
				temp.Next = temp2
				temp2 = temp2.Next
			}
			temp = temp.Next
		}
		if temp1 != nil {
			temp.Next = temp1
		} else if temp2 != nil {
			temp.Next = temp2
		}
		return dummyHead.Next
	}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

// ListNode 节点
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// Insert 向节点后插入
func (l *ListNode) Insert(val interface{}) *ListNode {
	n := &ListNode{val, l.Next}
	l.Next = n
	return l
}
