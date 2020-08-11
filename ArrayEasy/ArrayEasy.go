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

func merge(nums1 []int, m int, nums2 []int, n int) { // 88 合并两个有序数组
	point1 := 0 // 两个数组分别使用两个指针对元素进行比较 nums1clone指针
	point2 := 0

	nums1clone := make([]int, m) // 数组nums1深拷贝
	_ = copy(nums1clone, nums1[:m])

	for i := 0; i <= m+n; i++ {
		if point1 == m || point2 == n { // 每次循环前先判断 如果指针指到最后一个元素
			if point1 == m {
				_ = copy(nums1[i:m+n], nums2[point2:n])
			}
			if point2 == n {
				_ = copy(nums1[i:m+n], nums1clone[point1:m])
			}
			break
		} else if nums1clone[point1] <= nums2[point2] {
			nums1[i] = nums1clone[point1]
			point1++
		} else {
			nums1[i] = nums2[point2]
			point2++
		}
	}
}

func countCharacters(words []string, chars string) int { // 1106 拼写单词
	var length = 0 // 输出数组长度

	m := make(map[string]int) // 将chars保存至hash map
	for _, char := range chars {
		m[string(char)] = 0 // 对hash map初始化
	}
	for _, char := range chars {
		m[string(char)]++ // 记录每种char的个数
	}

	for _, word := range words {

		mclone := make(map[string]int)
		for k, v := range m { // map深拷贝
			mclone[k] = v
		}

		var canspell = true
	for1:
		for _, char := range word {
			if _, ok := mclone[string(char)]; ok { // 如果键存在
				mclone[string(char)]--
				if mclone[string(char)] < 0 {
					canspell = false
				}
			} else { // 如果键不存在
				canspell = false
				break for1
			}
		}
		if canspell == true { // 如果可以拼写单词
			length = length + len(word)
		}
	}
	return length
}

func average(salary []int) float64 { // 1491. 去掉最低工资和最高工资后的工资平均值
	var maxs = salary[0] // 最低工资
	var mins = salary[0] // 最高工资
	var sum = 0          // 工资之和
	for _, v := range salary {
		if v > maxs {
			maxs = v
		}
		if v < mins {
			mins = v
		}
		sum += v
	}
	return float64(sum-maxs-mins) / float64(len(salary)-2)
}

func main() {
	log.Println("ArrayEasy")

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

	// // 合并两个有序数组
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// nums2 := []int{2, 5, 6}
	// m := 3
	// n := 3
	// merge(nums1, m, nums2, n)
	// log.Println(nums1)

	// // 拼写单词
	// words1 := []string{"hello", "world", "leetcode"}
	// chars1 := "welldonehoneyr"

	// words2 := []string{"cat", "bt", "hat", "tree"}
	// chars2 := "atach"

	// k1 := countCharacters(words1, chars1)
	// k2 := countCharacters(words2, chars2)
	// log.Println(k1, k2)

	// // 去掉最低工资和最高工资后的工资平均值
	// salary1 := []int{8000, 9000, 2000, 3003, 6000, 1000}
	// salary2 := []int{6000, 5000, 4000, 3000, 2000, 1000}
	// k1 := average(salary1)
	// k2 := average(salary2)
	// log.Println(k1, k2)
}
