package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool { // 141.环形链表
	// 思路一：使用map来记录是否遍历同一个节点两次 时间复杂度O(n)
	// 思路二：时间复杂度为O(1)的算法
	var m = make(map[*ListNode]int) // 节点被遍历的次数
	pos := -1
	for head != nil {
		head = head.Next
	}
	if pos == -1 {
		return false
	}
	return true
}

func main() {
	log.Println("TwoPointersEasy")

	l3 := ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := ListNode{
		Val:  2,
		Next: &l3,
	}
	l1 := ListNode{
		Val:  1,
		Next: &l2,
	}

	log.Println(hasCycle(&l1)) // 141.环形链表
}
