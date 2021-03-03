package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode { // 86.分隔链表
	// 思路一：遍历所有节点找到第一个大于等于x的节点地址
	// bigger 存储第一个大于x的节点，preBiggeer存储bigger之前的那个节点，smallHead smallTail分别是小于x链表的头和尾
	var pre, preBigger, bigger, smallerHead, smallerTail *ListNode
	for head != nil {
		if head.Val > x && bigger == nil {
			bigger = head
			preBigger = pre
		} else {
			smaller.Next = head
		}
		pre = head
		head = head.Next
	}
	// 将小于x的节点链表接到bigger和头节点之间
	preBigger.Next = smallerHead
	smallerTail.Next = bigger

	// 返回结果
	return head
}

func main() {
	log.Println("TwoPointersMedium")

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
	log.Println(partition())
}
