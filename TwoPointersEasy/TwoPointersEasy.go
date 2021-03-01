package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleMap(head *ListNode) bool { // 141.环形链表
	// 思路一：使用map来记录是否遍历同一个节点两次 时间复杂度O(n)
	var m = make(map[*ListNode]int) // 节点被遍历的次数

	for head != nil {
		if m[head] == 2 {
			return true
		}
		m[head]++
		head = head.Next // 添加至map中
	}

	return false
}

func hasCycle(head *ListNode) bool { // 141.环形链表
	// 思路二：使用快慢指针，相遇则存在环形
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
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
	l3.Next = &l1

	log.Println(hasCycle(&l1)) // 141.环形链表
}
