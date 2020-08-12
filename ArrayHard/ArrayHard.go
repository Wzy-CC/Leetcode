package main

import "log"

func firstMissingPositive(nums []int) int { // 41.缺失的第一个正数 算法时间复杂度为O(n) 且只能使用常数级别的额外空间
	// 哈希碰撞法 大于len(nums)直接扔掉 负数和零直接扔掉 剩下的数字两两交换 len(nums)放置在nums[0]处
	for k,v := range nums{
		for 
	}
}

func numOfSubarraysNoContinuous(arr []int, k int, threshold int) int { // 1343 大小为k平均值大于阈值的子数组数目 子数组不连续
	return 1
}

func main() {
	log.Println("ArrayHard")

	// 缺失的第一个正数
	n := []int{100, 22, 3, 4, -1, 1}
	k := firstMissingPositive(n)
	log.Println(k)
}
