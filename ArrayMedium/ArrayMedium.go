package main

import (
	"log"
	"sort"
)

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

// // ?
// func movesToMakeZigzag(nums []int) int { // 1144. 递减元素使数组呈锯齿状
// 	var diffs []int // 差值数组
// 	var oddOpe = 0  // 计算对奇数索引操作的次数
// 	var ovenOpe = 0 // 计算对偶数索引操作的次数

// 	for i := 1; i < len(nums); i++ { // 每次计算相邻两个数的差值 当前减去前一个数的差值
// 		diff := nums[i] - nums[i-1]
// 		diffs = append(diffs, diff)
// 	}
// 	for k, v := range diffs { // 遍历差值数组 +-+- 形式
// 		if k%2 == 0 {

// 		}

// 	}
// 	if oddOpe <= ovenOpe { // 返回最小值
// 		return oddOpe
// 	}
// 	return ovenOpe
// }

func rotate(matrix [][]int) { // 面试题 旋转矩阵
	var tempMatrix = make([][]int, len(matrix)) // 行 空间申请
	for i := 0; i < len(matrix); i++ {
		tempMatrix[i] = make([]int, len(matrix)) // 列 空间申请
	}

	// 赋值到tempMetrix
	for rowindex, row := range matrix {
		for colindex, col := range row {
			tempMatrix[colindex][len(matrix)-rowindex-1] = col
		}
	}
	copy(matrix, tempMatrix) // 赋值
}

func getWinner(arr []int, k int) int { // 1535. 找出数组游戏的赢家
	// 数组本身由不同元素组成
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var wincount = 0 // 赢回合数
	var preIndex = 0 // 维护一个指针指向当前元素
	for {
		preIndex++
		// 每次比较preIndex和0号元素的值
		if arr[0] > arr[preIndex] { // 0号元素大
			wincount++
		} else { // 当前元素大
			arr[0] = arr[preIndex] // 赋值到第一个元素
			wincount = 1
		}

		if (wincount == min(k, len(arr))) || preIndex == len(arr)-1 {
			break
		}
	}
	return arr[0]
}

func maxDistance(position []int, m int) int { // 1552. 两球之间的磁力
	// 返回是否满足可以按照pred放入小球:贪心算法
	IfSatisfy := func(p []int, m int, pred int) bool {
		var lastindex = 0                    // 记录上一个放置小球的位置
		for i := 0; i < len(position); i++ { // 遍历position数组
			if i == 0 {
				m--
				lastindex = i
			}
			if p[lastindex]+pred <= p[i] { // 判断数组位置是否可以放置小球:可以放置
				m--
				lastindex = i
			}
		}
		if m > 0 { // 放置到最后一个位置还是有剩余，不行
			return false
		}
		return true
	}

	// 思路：先排序，理想情况为直接放置在等分索引处
	// for i := 0; i < len(position)-1; i++ { // 冒泡排序
	// 	for j := 0; j < len(position)-i-1; j++ {
	// 		if position[j] > position[j+1] {
	// 			position[j], position[j+1] = position[j+1], position[j]
	// 		}
	// 	}
	// }
	sort.Ints(position) // 注释:此处冒泡过不了
	// 计算理论上的最大化的最小距离:向零取整
	WishDistance := (position[len(position)-1] - position[0]) / (m - 1)
	// 二分查找，使用贪心算法分配小球位置
	var mindistance = 1            // 最小值从1计算起
	var maxdistance = WishDistance // 最大值从WishDistance起
	for {
		if mindistance == maxdistance {
			return 1
		}
		pred := (mindistance + maxdistance) / 2 // 二分查找当前的位置
		// 按照pred分配小球位置
		ifsatisfy := IfSatisfy(position, m, pred)
		if ifsatisfy { // 如果当前距离满足要求，左边界更新
			mindistance = pred
		} else { // 如果当前距离不满足要求，右边界更新
			maxdistance = pred
		}
		if mindistance == maxdistance-1 { // 终止条件:当min和max相差1
			if IfSatisfy(position, m, mindistance) && !IfSatisfy(position, m, maxdistance) { // 且 min符合要求 max不符合要求时
				return mindistance
			} else if IfSatisfy(position, m, mindistance) && IfSatisfy(position, m, maxdistance) { // 且 min符合要求 max符合要求时
				return maxdistance
			}
			break
		}
	}
	return 1
}

func minimumTotal(triangle [][]int) int { // 120. 三角形最小路径和
	// 动态规划: 降维 原地替换
	min := func(i ...int) int { // 定义min函数
		var m = i[0] // 初始化为第一个元素
		for _, v := range i {
			if v < m {
				m = v
			}
		}
		return m
	}

	var prerow = len(triangle) - 1 // 初始状态从最后一行遍历开始
	for {
		if prerow == 0 { // 如果当前行数为0
			return triangle[0][0] // 返回第一个元素
		}
		for i := 0; i < prerow; i++ { // 遍历prerow上一行
			triangle[prerow-1][i] = min((triangle[prerow-1][i] + triangle[prerow][i]), (triangle[prerow-1][i] + triangle[prerow][i+1])) // 替换
		}
		prerow = prerow - 1
		log.Println(triangle)
	}
}

// func pancakeSort(A []int) []int { // 969. 煎饼排序
// 	// 煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行），任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。
// 	var kseq []int // 返回k序列
// 	// check := func(a []int) bool { // 检查函数，检查A数组是否需要煎饼反转操作
// 	// 	for k,v := range a {
// 	// 		if
// 	// 	}
// 	// }
// 	reverse := func(a []int, k int) { // 定义反转操作:原地交换
// 		for i := 0; i < (k+1)/2; i++ { // k兼容奇数/偶数
// 			a[i], a[k-i-1] = a[k-i-1], a[i] // 交换
// 		}
// 	}
// 	// N次反转后保证最大元素位于序列右侧

// 	return kseq
// }

// func backtrace() {
// 	// for k,choice := range choices {
// 	// 	// 选择下一个元素
// 	// 	// 撤销选择
// 	// }
// 	back
// }
func subsets(nums []int) [][]int { // 78. 子集
	// 思路1:使用回溯算法
	// 思路2:递归每次递增子集添加到结果中
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var results = [][]int{}
	results = append(results, []int{}) // 首先添加空集
	for _, v := range nums {
		for _, result := range results { // 每添加一个数 需要将results中
			result = append(result, v)
			results = append(results, result) // 每次生成新的集合加入
			log.Println(">>>", len(result), cap(result))
			log.Println(results)
		}
	}
	return results
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

	// // 递减元素使数组呈锯齿状
	// nums1 := []int{9, 6, 1, 6, 2}
	// nums2 := []int{1, 2, 3}
	// k1 := movesToMakeZigzag(nums1)
	// k2 := movesToMakeZigzag(nums2)
	// log.Println(k1, k2)

	// // 旋转矩阵
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// rotate(matrix)
	// log.Println(matrix)

	// // 找出数组游戏的赢家
	// arr := []int{2, 1, 3, 5, 4, 6, 7}
	// k := 2
	// log.Println(getWinner(arr, k))

	// // 两球之间的磁力
	// position1 := []int{52, 63, 51, 8, 85, 27, 38, 30, 90, 66, 60, 6, 61, 72, 3, 81, 1, 21, 55, 13, 78, 94, 31, 96, 33, 40, 56, 77, 15, 79, 93, 82, 35, 2, 98, 28, 70, 5, 9, 88, 83, 14, 39, 54, 65, 74, 10, 48, 84, 92, 53, 71, 36, 20, 97, 25, 59, 50, 22, 12, 16, 11, 89, 86, 64}
	// m1 := 65
	// log.Println(maxDistance(position1, m1))

	// position2 := []int{1, 2, 3, 4, 7}
	// m2 := 3
	// log.Println(maxDistance(position2, m2))

	// // 三角形最小路径和
	// triangle := [][]int{
	// 	{2},
	// 	{3, 4},
	// 	{6, 5, 7},
	// 	{4, 1, 8, 3},
	// }
	// log.Println(minimumTotal(triangle))

	// // 煎饼排序
	// array := []int{3, 2, 4, 1}
	// log.Println(pancakeSort(array))

	// 子集
	nums := []int{1, 2, 3, 4, 5, 6}
	log.Println(subsets(nums))
}
