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

func searchInsert(nums []int, target int) int { // 35. 搜索插入位置 数组排序好
	var index int // 返回索引
	for k, v := range nums {
		if k == 0 && target < nums[k] {
			index = 0
			break
		} else if k == len(nums)-1 && target > nums[k] { // 最后一个元素
			index = len(nums)
			break
		} else if v < target && nums[k+1] > target { // 插入
			index = k + 1
			break
		} else if v == target { // 查找到target存在数组内
			index = k
			break
		}
	}
	return index
}

func replaceElements(arr []int) []int { // 1299. 将每个元素替换为右侧最大元素
	for i := 0; i < len(arr); i++ { // 每次替换一个数字
		var tempmax int      // 右侧最大数字
		if i == len(arr)-1 { // 如果是最后一个数字
			arr[i] = -1
			break
		}
		tempmax = arr[i+1] // tempmax 值初始默认为i+1
		for j := i + 1; j < len(arr); j++ {
			if tempmax < arr[j] {
				tempmax = arr[j]
			}
		}
		arr[i] = tempmax
	}
	return arr
}

func generate(numRows int) [][]int { // 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行
	var triangle [][]int
	var lastRow []int // 上一排数组
	for i := 1; i <= numRows; i++ {
		var tempRow = make([]int, i) // 每一排
		tempRow[0] = 1               // 第一个元素赋值
		tempRow[i-1] = 1             // 最后一个元素赋值
		for j := 1; j < i-1; j++ {   // 中间元素赋值
			tempRow[j] = lastRow[j] + lastRow[j-1]
		}
		triangle = append(triangle, tempRow)
		lastRow = make([]int, i)
		_ = copy(lastRow, tempRow)
	}
	return triangle
}

func twoSum(numbers []int, target int) []int { // 167. 两数之和 II - 输入有序数组
	var results = make([]int, 2) // 返回结果
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] > target { // 已经超过了
				break
			} else if numbers[i]+numbers[j] == target { // 返回结果
				results[0] = i + 1
				results[1] = j + 1
				break
			}
		}
	}
	return results
}

func minTimeToVisitAllPoints(points [][]int) int { // 1266. 访问所有点的最小时间
	var routeTime = 0
	var tempTime int
	var lastpoint = make([]int, 2) // 上一个节点
	for k, point := range points { // 遍历节点
		if k == 0 { // 第一个节点
			_ = copy(lastpoint, point)
		} else {
			xmove := point[0] - lastpoint[0]
			if xmove < 0 {
				xmove = -xmove
			}
			ymove := point[1] - lastpoint[1]
			if ymove < 0 {
				ymove = -ymove
			}
			if diff := ymove - xmove; diff < 0 { // xmove比较大
				tempTime = xmove
			} else {
				tempTime = ymove
			}
			routeTime += tempTime
			_ = copy(lastpoint, point)
		}
	}
	return routeTime
}

func numPairsDivisibleBy60(time []int) int { // 1010. 总持续时间可被 60 整除的歌曲 超长输入
	// 本题容易做成O(N^2)的时间复杂度,时间复杂度为O(N)为正确做法
	var songNum = 0 // 返回被60整除的歌曲数量
	var hashMap = make(map[int]int)
	for _, v := range time {
		hashMap[v%60]++
	}
	for i := 1; i < 30; i++ { // 从索引1开始计算 额外关心歌曲时间长度是60倍数的时长和余数为30的时长
		multip := hashMap[i] * hashMap[60-i]
		songNum = songNum + multip
	}
	for i := 0; i < hashMap[0]; i++ { // 额外计算0
		songNum = songNum + i
	}
	for i := 0; i < hashMap[30]; i++ { // 额外计算30
		songNum = songNum + i
	}
	return songNum
}

func duplicateZeros(arr []int) { // 1089. 复写零
	var temparr []int       // 临时切片用于复制
	for _, v := range arr { // 遍历找0
		if v == 0 {
			temparr = append(temparr, v)
			temparr = append(temparr, 0)
		} else {
			temparr = append(temparr, v)
		}
	}
	copy(arr, temparr)
	// arr = temparr[:] // 截断赋值不行，slice知识点详情见https://juejin.im/user/1292681406870013
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

	// // 搜索插入位置
	// nums := []int{1}
	// target := 1
	// k := searchInsert(nums, target)
	// log.Println(k)

	// // 将每个元素替换为右侧最大元素
	// arr := []int{17, 18, 5, 4, 6, 1}
	// k := replaceElements(arr)
	// log.Println(k)

	// // 杨辉三角
	// numRows := 7
	// k := generate(numRows)
	// log.Println(k)

	// // 两数之和 II
	// numbers := []int{2, 7, 11, 15}
	// target := 9
	// k := twoSum(numbers, target)
	// log.Println(k)

	// // 访问所有点的最小时间
	// points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	// k := minTimeToVisitAllPoints(points)
	// log.Println(k)

	// // 1010. 总持续时间可被 60 整除的歌曲
	// time := []int{30, 20, 150, 100, 40}
	// k := numPairsDivisibleBy60(time)
	// log.Println(k)

	// // 复写零
	// arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	// duplicateZeros(arr)
	// log.Println(arr)
}
