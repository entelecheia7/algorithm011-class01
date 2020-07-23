package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTotal2([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

// 最长公共子序列题目
// 三角形最小路径和
// 法一：动态规划，O(n^2)
func minimumTotal1(triangle [][]int) (minPath int) {
	level := len(triangle)
	if level == 0 || len(triangle[0]) == 0 {
		return 0
	}
	if level == 1 {
		return triangle[0][0]
	}
	state := make([][]int, level)
	for i := 0; i < level; i++ {
		state[i] = make([]int, i+1)
	}
	state[0][0] = triangle[0][0]
	for i := 1; i < level; i++ { // 计算每行的路径
		for j := 0; j <= i; j++ {
			state[i][j] = triangle[i][j]
			if i == j {
				state[i][j] += state[i-1][j-1]
			} else if j == 0 {
				state[i][j] += state[i-1][j]
			} else {
				state[i][j] += getMin(state[i-1][j-1], state[i-1][j])
			}
		}
	}
	minPath = math.MaxInt64
	for _, path := range state[level-1] {
		minPath = getMin(minPath, path)
	}
	return minPath
}

// 法二：对法一进行空间优化，只使用一个一维数组保存中间状态
// 同时自底向上计算，省去正向计算最后查找最小路径的一次循环
// 时间O(n^2)，空间O(n)，n是triangle行数
// best
func minimumTotal2(triangle [][]int) int {
	level := len(triangle)
	if level == 0 || len(triangle[0]) == 0 {
		return 0
	}
	if level == 1 {
		return triangle[0][0]
	}
	state := make([]int, level+1)
	for i := level - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			state[j] = triangle[i][j] + getMin(state[j], state[j+1])
		}
	}

	return state[0]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 最大子序和
// 法一：动态规划
func maxSubArray1(nums []int) (max int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max = nums[0]
	pre := nums[0]
	for i := 1; i < n; i++ {
		pre = getMax(nums[i], pre+nums[i])
		max = getMax(max, pre)
	}
	return max
}

// 法二：分治
// O(nlogn)
func maxSubArray2(nums []int) (max int) {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	mid := n >> 1
	maxLeft := maxSubArray2(nums[0:mid])
	maxRight := maxSubArray2(nums[mid:])
	// 计算中间向两边最大子序和
	l := nums[mid-1] // 向左侧最大值
	tmp := 0
	for i := mid - 1; i >= 0; i-- {
		tmp += nums[i]
		l = getMax(tmp, l)
	}
	r := nums[mid] // 向右侧最大值
	tmp = 0
	for i := mid; i < n; i++ {
		tmp += nums[i]
		r = getMax(tmp, r)
	}

	return getMax(getMax(maxLeft, maxRight), l+r)
}

// 打家劫舍
// 法一：动态规划。
// nums[i] 要么纳入统计，要么不纳入统计。
// 计算两种情况，取较大值
// f(2) = getMax(nums[1]+f(0), f(1))
// f(3) = getMax(nums[2]+f(1), f(2))
// 递推公式：f(n) = getMax(nums[n-1]+ f(n-2), f(n-1))
// 空间、时间O(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	state := make([]int, n+1) // state[i]表示前i个元素可以获取的最大结果
	state[0] = 0
	state[1] = nums[0]
	for i := 2; i <= n; i++ {
		state[i] = getMax(nums[i-1]+state[i-2], state[i-1])
	}

	return state[n]
}

// 法二：对法一的递推进行改进，节省空间
func rob2(nums []int) (result int) {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	pp, p := 0, nums[0]
	for i := 2; i <= n; i++ {
		result = getMax(nums[i-1]+pp, p)
		pp, p = p, result
	}
	return result
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
