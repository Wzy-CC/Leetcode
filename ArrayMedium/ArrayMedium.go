package main

import "log"

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

func tictactoe(board []string) string { // 面试题16.04 井字游戏
	var spaceFlag = 0              // 0 default none space 1 exist space
	stringToNum := map[string]int{ // X:1 O:0 " ":-1000
		"X": 1,
		"O": 0,
		" ": -1000,
	}

	for i := 0; i < len(board); i++ { // 使用双层for循环来计算列和 列号
		var sum = 0                       // 列和初始化
		for j := 0; j < len(board); j++ { // 行号
			sum += stringToNum[string(board[j][i])]
		}
		if sum == len(board) { // winner : X
			return "X"
		} else if sum == 0 { // winner : O
			return "O"
		} else if sum < 0 {
			spaceFlag = 1
		}
	}

	for i := 0; i < len(board); i++ { // 使用双层for循环来计算列和 行号
		var sum = 0                       // 行和初始化
		for j := 0; j < len(board); j++ { // 列号
			sum += stringToNum[string(board[i][j])]
		}
		if sum == len(board) { // winner : X
			return "X"
		} else if sum == 0 { // winner : O
			return "O"
		} else if sum < 0 {
			spaceFlag = 1
		}
	}

	var sum1 = 0                      // 右上左下对角线之和
	var sum2 = 0                      // 左上右下对角线之和
	for i := 0; i < len(board); i++ { // 计算对角线之和
		sum1 = sum1 + stringToNum[string(board[len(board)-i-1][i])]
		sum2 = sum2 + stringToNum[string(board[i][i])]
	}
	if sum1 == len(board) || sum2 == len(board) { // winner : X
		return "X"
	} else if sum1 == 0 || sum2 == 0 { // winner : O
		return "O"
	} else if sum1 < 0 || sum2 < 0 {
		spaceFlag = 1
	}

	if spaceFlag == 1 { // 如果不存在判断是否有空格
		return "Pending"
	} else {
		return "Draw"
	}
}

func merge1(intervals [][]int) [][]int { // 56. Merge Intervals 合并区间
	var results [][]int // 返回结果
	var m map[int]int   // key int值 value 有效位标识是否为边界数字 1:右边界数字 0:非边界数字 -1:左边界数字
	m = make(map[int]int)
	// step1:区间去重
	for _, interval := range intervals { // 遍历外层
		for i := interval[0]; i <= interval[1]; i++ { // 向map添加元素
			if i == interval[0] {
				m[i] = -1
			} else if i == interval[1] {
				m[i] = 1
			} else {
				m[i] = 0
			}
		}
	}
	// step2:排序
	var intervalSortSet []int
	for k, _ := range m {
		intervalSortSet = append(intervalSortSet, k)
	}
	for i := 0; i < len(intervalSortSet); i++ { // 冒泡排序
		for j := 0; j < len(intervalSortSet)-i-1; j++ {
			if intervalSortSet[j] > intervalSortSet[j+1] {
				intervalSortSet[j], intervalSortSet[j+1] = intervalSortSet[j+1], intervalSortSet[j]
			}
		}
	}
	// step3:隔断区间
	var startindex int // 区间起始
	var endindex int   // 区间终止
	startindex = 0
	for k, v := range intervalSortSet {
		if k == len(intervalSortSet)-1 { // k是最后一个数
			endindex = k
			var tempinterval = make([]int, 2)
			tempinterval[0] = intervalSortSet[startindex]
			tempinterval[1] = intervalSortSet[endindex]
			results = append(results, tempinterval[:]) // 区间添加 重新开始计算间隔
		} else {
			if (m[v] == -1 && m[intervalSortSet[k+1]] == 1 && v == intervalSortSet[k+1]-1) || v != intervalSortSet[k+1]-1 { // 如果当前跳动范围大于1 或者连续但是当前为左边界 下一个为右边界 隔断产生
				endindex = k
				var tempinterval = make([]int, 2)
				tempinterval[0] = intervalSortSet[startindex]
				tempinterval[1] = intervalSortSet[endindex]
				results = append(results, tempinterval) // 区间添加 重新开始计算间隔
				startindex = k + 1
			}
		}
	}
	return results
}

func merge(intervals [][]int) [][]int { // 56. Merge Intervals 合并区间
	if len(intervals) == 0 {
		return [][]int{}
	}
	var results [][]int // 返回结果
	// step1 对区间进行排序 按照左端点大小进行排序
	for i := 0; i < len(intervals); i++ { // 冒泡排序
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] { // 左端点按照从小到大排序
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	// step2 对区间进行合并
	var tempinterval []int // 临时区间，用于更新右端点
	tempinterval = make([]int, 2)
	_ = copy(tempinterval, intervals[0])
	for k, interval := range intervals {
		if interval[0] > tempinterval[1] { // 直到左端点大于右端点为止
			results = append(results, tempinterval)
			tempinterval = interval
			if k == len(intervals)-1 { // 如果到达最后一个节点
				results = append(results, tempinterval)
			}
		} else { // 更新右端点
			if tempinterval[1] < interval[1] {
				tempinterval[1] = interval[1]
			}
			if k == len(intervals)-1 { // 如果到达最后一个区间
				results = append(results, tempinterval)
			}
		}
	}
	return results
}

func movesToMakeZigzag(nums []int) int { // 1144. 递减元素使数组呈锯齿状
	var diffs []int // 差值数组
	var oddOpe = 0  // 计算对奇数索引操作的次数
	var ovenOpe = 0 // 计算对偶数索引操作的次数

	for i := 1; i < len(nums); i++ { // 每次计算相邻两个数的差值 当前减去前一个数的差值
		diff := nums[i] - nums[i-1]
		diffs = append(diffs, diff)
	}
	for k, v := range diffs { // 遍历差值数组 +-+- 形式
		if k%2 == 0 {

		}

	}
	if oddOpe <= ovenOpe { // 返回最小值
		return oddOpe
	}
	return ovenOpe
}

func main() {
	log.Println("ArrayMedium")

	// // 大小为k平均值大于阈值的子数组数目
	// arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	// k := numOfSubarrays(arr, 3, 4)
	// log.Println(k)

	// // 面试题 16.04 井字游戏
	// board := []string{"O"}
	// k := tictactoe(board)
	// log.Println(k)

	// // 56. 合并区间
	// interval := [][]int{
	// 	[]int{1, 3},
	// 	[]int{2, 6},
	// 	[]int{8, 10},
	// 	[]int{15, 18},
	// }
	// k := merge(interval)
	// log.Println(k)

	// 递减元素使数组呈锯齿状
	nums1 := []int{9, 6, 1, 6, 2}
	nums2 := []int{1, 2, 3}
	k1 := movesToMakeZigzag(nums1)
	k2 := movesToMakeZigzag(nums2)
	log.Println(k1, k2)
}
