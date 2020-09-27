package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode { // 2. 两数相加
	var startFlag = 1 // 开始标志,1起始
	var endFlag = 0   // 结束标志,当两个链表遍历完成时退出循环

	var lstart *ListNode   // 结果指针
	var tempNext *ListNode // 指向下一个节点的指针
	var ltemp ListNode     // 结果单链表 当前节点
	var carryFlag = 0      // 进位标志 0无进位 1存在进位
	var sum int            // 结果Val

	for {
		sumtemp := l1.Val + l2.Val + carryFlag // 循环遍历两个单链表,将两个元素和进位标志相加
		if sumtemp > 9 {                       // 存在进位
			carryFlag = 1
			sum = sumtemp - 10
		} else {
			carryFlag = 0
			sum = sumtemp // 如果不存在进位直接赋值
		}

		ltemp.Val = sum
		ltemp.Next = tempNext // 当前为野指针

		if startFlag == 1 { // 判断当前是否为第一个元素
			lstart = &ltemp
			startFlag = 0 // 置为无效
		}

		lresult.Val = sum

		if l1.Next == nil { // 判断一下单链表是否到达尽头
		}
		if l2.Next == nil {
		}
		l1 = l1.Next
		l2 = l2.Next

		if endFlag > 1 { // endFlag为2时表示两个单链表都遍历完成
			break
		}
	}
	return &lstart
}

func main() {
	log.Println("LinkedList")

	// 2. 两数相加
	// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	// 输出：7 -> 0 -> 8
	l13 := ListNode{
		Val:  3,
		Next: nil,
	}
	l12 := ListNode{
		Val:  4,
		Next: &l13,
	}
	l11 := ListNode{
		Val:  2,
		Next: &l12,
	}

	l23 := ListNode{
		Val:  4,
		Next: nil,
	}
	l22 := ListNode{
		Val:  6,
		Next: &l23,
	}
	l21 := ListNode{
		Val:  5,
		Next: &l22,
	}
	lsum := addTwoNumbers(&l11, &l21)
	lt := lsum // 起始lt
	for {
		log.Println(lt.Val)
		if lt.Next == nil {
			break
		}
		lt = lt.Next
	}
}
