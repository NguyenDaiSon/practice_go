package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) PrintList() {
	fmt.Printf("[ ")
	for it := head.Next; it != nil; it = it.Next {
		fmt.Printf("%d ", it.Val)
	}
	fmt.Printf("]\n")
}

func (head *ListNode) BuildList(values []int) *ListNode {
	var tail *ListNode
	for _, value := range values {
		newTail := &ListNode{Val: value}
		if head.Next == nil {
			head.Next = newTail
		} else {
			tail.Next = newTail
		}
		tail = newTail
	}
	return head
}

func (head *ListNode) InsertSort() *ListNode {
	if head.Next == nil {
		return head
	}

	for it := head.Next.Next; it != nil; it = it.Next {
		for jt := head.Next; jt != it; jt = jt.Next {
			if it.Val < jt.Val {
				it.Val, jt.Val = jt.Val, it.Val
			}
		}
	}

	return head
}

func main() {
	head := &ListNode{}

	head.BuildList([]int{2, 1, 3, 4}).PrintList()

	///////////////////////////////////////////
	// LC 147. Insertion Sort List
	///////////////////////////////////////////
	head.InsertSort().PrintList()
}
