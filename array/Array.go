package main

import (
	// "github.com/getlantern/deepcopy"
	"log"
	"sort"
)

func heightChecker(heights []int) int {
	var count int = 0
	// var heightsclone []int
	var heightsclone = make([]int, len(heights[:7:7]))
	_ = copy(heightsclone, heights[:])
	sort.Ints(heights)
	log.Printf("%p\n", heightsclone)
	log.Printf("%p\n", heights)
	log.Println(heightsclone)
	log.Println(heights)
	for i := 0; i < len(heights); i++ {
		if heights[i] != heightsclone[i] {
			count = count + 1
		}
	}
	return count
}

func missingNumber(nums []int) int { // 剑指offer53
	for k, v := range nums { // 中间缺少数字
		if k != v {
			return k
		}
	}
	return len(nums) // 两边缺少数字
}

func findLengthOfLCIS(nums []int) int { // 674 最长连续递增序列
	if len(nums) == 0 {
		return 0
	}
	var startindex int // 递增子序列开始的索引
	var maxlength int  // 最大递增子序列的长度
	maxlength = 1
	for k, v := range nums {
		if k == 0 {
			startindex = k
			continue
		} else if v <= nums[k-1] || k == len(nums)-1 { // 如果小于等于前面一个数字 或者到达最后一个元素时
			if k == len(nums)-1 && v > nums[k-1] {
				k++
			}
			if maxlength < k-startindex {
				maxlength = k - startindex
			}
			startindex = k
		}
	}
	return maxlength
}

func removeDuplicates(nums []int) int { // 26 删除排序数组中的重复项 原地修改数组 空间复杂度O(1)
	var numsTomap map[int]int
	for k, v := range nums {
		nums[v] = 1
	}

	maplength := len(numsTomap)
	var count int // 控制索引
	for k,v := range  { // map重新赋值给nums
		nums[count] = 
	}
	return maplength
}

func main() {
	// 高度检查
	// h := []int{34, 1, 2, 3, 4, 5, 7}
	// heightChecker(h)

	// // 剑指offer53
	// k := missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9})
	// log.Println(k)

	// // 最长连续递增序列
	// k := findLengthOfLCIS([]int{2})
	// log.Println(k)

	// 删除排序数组中的重复项
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	k := removeDuplicates(n)
	log.Println(k, "\n")
	log.Println(n[0])
}
