package main

import (
	"log"
	"sort"
)

func reversePairs(nums []int) int { // 剑指 Offer 51. 数组中的逆序对
	var res int
	// def func
	var merge func([]int, []int) []int
	var mergeSort func([]int) []int
	// func body
	merge = func(l, r []int) []int { // solve
		var lpoint = len(l) - 1 // l index
		var rpoint = len(r) - 1 // r index
		var res []int           // result stack
		for lpoint >= 0 && rpoint >= 0 {
			if l[lpoint] < r[rpoint] {
				res = res + len(r)
			}
		}
		res = append(res, l[:lpoint]...)
		res = append(res, r[:rpoint]...)
		return res
	}
	mergeSort = func(a []int) []int { // devide
		if len(a) <= 1 {
			return a
		}
		mid := len(a) / 2
		left := mergeSort(a[:mid])
		right := mergeSort(a[mid:])
		return merge(left, right)
	}
	return res
}

func main() {
	log.Println("剑指offer")

	// 数组中的逆序对
	arr := []int{7, 5, 6, 4}
	log.Println(reversePairs(arr))
}
