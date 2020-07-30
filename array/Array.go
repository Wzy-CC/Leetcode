package main

import "log"

func heightChecker(heights []int) int { // 高度检查器
	var heightsclone = make([]int, len(heights[:])) // 克隆原数组
	_ = copy(heightsclone, heights[:])

	for i := 0; i < len(heights); i++ { // 原地冒泡排序
		for j := 0; j < len(heights)-i-1; j++ {
			if heights[j] > heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
			}
		}
	}

	var count int = 0
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

func removeDuplicates(nums []int) int { // 26 删除排序数组中的重复项 给定排序好数组 原地修改数组 空间复杂度O(1)
	var point = 0 // 指针 赋值数组
	for k, v := range nums {
		if k == 0 { // 判断是否为第一个元素
			point++
		} else if nums[k-1] != nums[k] { // 当前元素和前一个元素不相同
			nums[point] = v
			point++
		}
	}
	return point
}

func numOfSubarrays(arr []int, k int, threshold int) int { // 1343 大小为k平均值大于阈值的子数组数目 子数组连续
	var count = 0            // 起始个数为0个
	var sum = 0              // 数组之和
	for i := 0; i < k; i++ { // 计算起始sum
		sum += arr[i]
	}
	if float64(sum)/float64(k) >= float64(threshold) {
		count++
	}

	for i := 1; i < len(arr)-k+1; i++ { // 滑动窗口
		sum = sum - arr[i-1] + arr[i+k-1]
		if float64(sum)/float64(k) >= float64(threshold) {
			count++
		}
	}
	return count
}

func numOfSubarraysNoContinuous(arr []int, k int, threshold int) int { // 1343 大小为k平均值大于阈值的子数组数目 子数组不连续
	return 1
}

func main() {
	// // 高度检查
	// h := []int{34, 1, 2, 3, 4, 5, 7}
	// log.Println(heightChecker(h))

	// // 剑指offer53
	// k := missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9})
	// log.Println(k)

	// // 最长连续递增序列
	// k := findLengthOfLCIS([]int{2})
	// log.Println(k)

	// // 删除排序数组中的重复项
	// n := []int{0, 1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4, 5}
	// k := removeDuplicates(n)
	// log.Println(k, "\n")
	// log.Println(n)

	// 大小为k平均值大于阈值的子数组数目
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k := numOfSubarrays(arr, 3, 4)
	log.Println(k)

}
