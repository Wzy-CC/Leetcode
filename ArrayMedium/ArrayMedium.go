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

func movesToMakeZigzag(nums []int) int { // 1144. 递减元素使数组呈锯齿状

}

func main() {

	// // 大小为k平均值大于阈值的子数组数目
	// arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	// k := numOfSubarrays(arr, 3, 4)
	// log.Println(k)

	// // 面试题 16.04 井字游戏
	// board := []string{"O"}
	// k := tictactoe(board)
	// log.Println(k)

	// 1144. 递减元素使数组呈锯齿状
	k := movesToMakeZigzag()
	log.Println(k)

	log.Println("l")
}
