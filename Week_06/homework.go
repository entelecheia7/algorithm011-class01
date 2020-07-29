package main

import (
	"fmt"
	"sort"
)

func main() {
	// 最小路径和

	// 解码方法
	// fmt.Println(numDecodings("0"))    // 0
	// fmt.Println(numDecodings("0001")) // 0
	// fmt.Println(numDecodings("12"))   // 2
	// fmt.Println(numDecodings("27"))   // 1

	// 最大正方形
	// fmt.Println(maximalSquare([][]byte{
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'0', '0', '0', '0', '0'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'}})) // 4

	// 任务调度器
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)) // 8

	// 回文子串

	// 最长有效括号

	// 编辑距离

	// 矩形区域不超过 K 的最大数值和

	// 青蛙过河

	// 分割数组的最大值

	// 学生出勤记录 II

	// 最小覆盖子串

	// 戳气球

}

/* 中等 */

// 最小路径和（亚马逊、高盛集团、谷歌在半年内面试中考过）

// 解码方法（亚马逊、Facebook、字节跳动在半年内面试中考过）
func numDecodings(s string) (cur int) {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	} else if n == 1 {
		return 1
	}
	p, pp := 1, 1
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			// 出现'0'有两种情况，一种是10或20，一种是当前的'0'不合法
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			cur = pp
		} else if (s[i-1] == '1') || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = p + pp
		} else {
			cur = p
		}
		p, pp = cur, p
	}
	return cur
}

// 最大正方形（华为、谷歌、字节跳动在半年内面试中考过）
func maximalSquare(matrix [][]byte) int {
	// check
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	maxBorder := 0
	dp := make([]int, n)
	var leftup int
	for i := 0; i < m; i++ {
		leftup = 0
		for j := 0; j < n; j++ {
			nextLeftup := dp[j]
			if matrix[i][j] == '1' {
				if maxBorder == 0 {
					maxBorder = 1
				}
				if i == 0 || j == 0 {
					dp[j] = 1
				} else {
					dp[j] = getMin(leftup, getMin(dp[j], dp[j-1])) + 1
				}
				if dp[j] > maxBorder {
					maxBorder = dp[j]
				}
			} else {
				dp[j] = 0
			}

			leftup = nextLeftup
		}
	}

	return maxBorder * maxBorder
}

// 任务调度器（Facebook 在半年内面试中常考）
func leastInterval(tasks []byte, n int) (time int) {
	if len(tasks) <= 1 || n == 0 {
		return len(tasks)
	}
	var task taskList
	count := 0
	for _, t := range tasks {
		task[t-'A'].task = t
		task[t-'A'].num++
		count++
	}
	// 按照任务数量从大到小排序
	sort.Sort(&task)
	// 完成数量最多的任务至少需要的时间
	time = (task[0].num-1)*(n+1) + 1
	// 如果有任务数量和最多数量相等的任务，那就需要+1
	i := 1
	for i < 26 && task[i].num == task[0].num {
		time++
		i++
	}
	// 如果计算结果小于数组数量，说明任务的种类和间隔时间要求对任务调度没有影响。
	if time < len(tasks) {
		return len(tasks)
	}
	return time
}

type taskList [26]taskData
type taskData struct {
	task byte
	num  int
}

func (t taskList) Len() int {
	return 26
}
func (t taskList) Less(i, j int) bool {
	return t[i].num > t[j].num
}
func (t *taskList) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

// 回文子串（Facebook、苹果、字节跳动在半年内面试中考过）

/* 困难 */

// 最长有效括号（字节跳动、亚马逊、微软在半年内面试中考过）
// 法一：动态规划
func longestValidParentheses1(s string) (maxSubLen int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i >= 2 && dp[i-2] > 0 { // 处理()()
					dp[i] += dp[i-2]
				}
			} else if s[i-1] == ')' { // 处理(())、()(())
				if dp[i-1] > 0 && i-dp[i-1] >= 1 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1] >= 2 && dp[i-dp[i-1]-2] > 0 {
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
		}
		if dp[i] > maxSubLen {
			maxSubLen = dp[i]
		}
	}

	return maxSubLen
}

// 法二：计数法
// 对左括号和右括号数量进行统计，当个数相等时，记录数目
// 当右括号数量大于左括号时，重新开始一轮统计
// 为了处理 (() 这种情况，再从右向左统计一遍
func longestValidParentheses2(s string) (maxLen int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	left, right := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
			if left == right {
				maxLen = getMax(maxLen, left+right)
			} else if right > left {
				left, right = 0, 0
			}
		}

	}
	left, right = 0, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
			if left == right {
				maxLen = getMax(maxLen, left+right)
			} else if left > right {
				left, right = 0, 0
			}
		} else {
			right++
		}
	}

	return maxLen
}

// 法三：栈解法
// 左括号直接入栈
// 遇右括号弹出栈顶元素，计算长度；如果栈为空，将当前位置入栈
func longestValidParentheses3(s string) (maxLen int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	stack := []int{-1} // 为了计算方便括号的长度
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = getMax(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

// 编辑距离（字节跳动、亚马逊、谷歌在半年内面试中考过）

// 矩形区域不超过 K 的最大数值和（谷歌在半年内面试中考过）

// 青蛙过河（亚马逊、苹果、字节跳动在半年内面试中考过）

// 分割数组的最大值（谷歌、亚马逊、Facebook 在半年内面试中考过）

// 学生出勤记录 II （谷歌在半年内面试中考过）

// 最小覆盖子串（Facebook 在半年内面试中常考）

// 戳气球（亚马逊在半年内面试中考过）

/* helper */
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
