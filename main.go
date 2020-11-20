package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 7, 8}
	head := fromArr(arr)
	l, r := head, head.Next
	fmt.Println(l, r)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func fromArr(arr []int) *ListNode {
	head := new(ListNode)
	cur := head
	for _, v := range arr {
		cur.Val = v
		cur.Next = new(ListNode)
		cur = cur.Next
	}
	return head
}

func (l *ListNode) toArr() (arr []int) {
	for l.Next != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return
}

func swapNext(l, r *ListNode) {
	t := l.Next.Next
	l.Next.Next = r.Next.Next
	r.Next.Next = t
	t = r.Next
	r.Next = l.Next
	l.Next = t
}
