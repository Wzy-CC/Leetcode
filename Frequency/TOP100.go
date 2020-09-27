package main

import (
	"log"
	"sort"
)

type TreeNode struct { // 二叉树
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode { // 236. 二叉树的最近公共祖先
	if root == nil { // 此时不可能查询到结果
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val { // 子树中寻找到结果节点，返回root
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 寻找左子树
	right := lowestCommonAncestor(root.Right, p, q) // 寻找右子树

	if left == nil { // 从下一层返回来的查询结果为nil 没有找到
		return right
	} else if right == nil { // 从下一层返回来的查询结果为nil 没有找到
		return left
	} else { // 当左右子树都找到时返回root
		return root
	}

	return nil // 当在此棵子树上进行查找无pq时，返回nil
}

func search(nums []int, target int) int { // 33. 搜索旋转排序数组
	// 	假设按照升序排序的数组在预先未知的某个点上进行了旋转。
	// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
	// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
	// 你可以假设数组中不存在重复的元素。
	// 你的算法时间复杂度必须是 O(log n) 级别。

	// 思路:二分搜索 使用mid判断左有序 还是 右有序
	if len(nums) == 0 {
		return -1
	}

	var left, right, mid int // 数组索引
	left = 0                 // 初始化
	right = len(nums) - 1

	for left <= right {
		mid = (left + right + 1) / 2 // 计算当前mid
		if nums[mid] == target {     // 命中
			return mid
		}
		if ((target <= nums[mid]) && (target >= nums[left])) || (nums[mid] <= nums[right] && (target > nums[right] || target < nums[mid])) { // 如果左有序 且target存在左边
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func coinChange(coins []int, amount int) int { // 322
	dp := make([]int, amount+1) // dp 记录问题的解
	for i := 1; i <= amount; i++ {
		dp[i] = -1                // dp初始化
		for _, c := range coins { // 解决问题规模为i的问题
			if i < c || dp[i-c] == -1 {
				continue
			}
			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}

func threeSum(nums []int) [][]int { // 15. 三数之和
	// leetcode answer
	var results [][]int // result
	sort.Ints(nums)
	var fixed, point1, point2 int // index
	for i := 0; i < len(nums)-2; i++ {
		// judge the val is same as preindex
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// init: fixed point1 point2
		fixed = i
		point1 = i + 1
		point2 = len(nums) - 1
		// left/right left
		for point1 < point2 {
			// calcul sum
			sum := nums[fixed] + nums[point1] + nums[point2]
			if sum == 0 {
				result := []int{nums[fixed], nums[point1], nums[point2]}
				results = append(results, result)
				for point1 < point2 && nums[point1] == nums[point1+1] {
					point1++
				}
				for point1 < point2 && nums[point2] == nums[point2-1] {
					point2--
				}
				point1++
				point2--
			} else { // adjust point1 or point2
				if sum > 0 { // right left
					point2--
				} else if sum < 0 { // left left
					point1++
				}
			}
		}
	}
	return results
}

func inorderTraversal(root *TreeNode) []int { // 94. 二叉树的中序遍历
	// 递归 & 迭代
	// 递归
	var results []int            // last results
	var dfs func(*TreeNode)      // def dfs
	dfs = func(root *TreeNode) { // 二叉树的中序遍历
		if root.Left == nil || root.Right == nil { // 如果遍历到叶子节点
			results = append(results, root.Val)
			return
		}
		results = append(results, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return results
}

func solveNQueens(n int) [][]string { // 面试题 08.12. 八皇后
	var traceback func() // 定义回溯函数
	traceback = func() {
		return
	}
	// 做出选择

	// 递归

	// 撤销选择
}

func main() {
	log.Println("Frequency")

	// 二叉树的最近公共祖先
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	log.Println(lowestCommonAncestor(root, root.Left.Left, root.Right.Left).Val)

	// 搜索旋转排序数组
	nums := []int{1, 3, 5}
	log.Println(search(nums, 5))

	// 零钱兑换
	coins := []int{1, 2, 5}
	log.Println(coinChange(coins, 11))

	// 三数之和
	nums1 := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	log.Println(threeSum(nums1))

	// 二叉树的中序遍历
	log.Println(inorderTraversal(root))

	// N皇后问题
	log.Println(solveNQueens(4))
}
