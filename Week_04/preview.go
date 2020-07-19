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
// 打家劫舍
