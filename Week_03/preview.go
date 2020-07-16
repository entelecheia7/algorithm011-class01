package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的层次遍历
// 法一：使用队列
func levelOrder1(root *TreeNode) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelNum := len(queue)
		levelVal := make([]int, levelNum)
		for i := 0; i < levelNum; i++ {
			levelVal[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelNum:]
		result = append(result, levelVal)
	}
	return result
}

// 法二：DFS
func levelOrder2(root *TreeNode) (result [][]int) {
	levelOrderDFSHelper(root, 0, &result)
	return result
}

// level表示节点层级，由上至下从0开始递增
func levelOrderDFSHelper(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if level == len(*result) {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	levelOrderDFSHelper(root.Left, level+1, result)
	levelOrderDFSHelper(root.Right, level+1, result)
}

// 分发饼干
// 贪心算法，O(nlogn)
func findContentChildren(g []int, s []int) (satisfy int) {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	biscuit := 0 // 代表饼干的下标
	child := 0   // 代表小孩的下标
	for biscuit < len(s) && child < len(g) {
		if s[biscuit] >= g[child] {
			child++
			satisfy++
		}
		biscuit++
	}
	return satisfy
}

// 买卖股票的最佳时机 II
// 这道题有一个误区：
// 不应该寻找差值最大的两个点买入卖出，而要寻找价差频繁交易.
// 最大价差策略无法处理回撤再上升的情况。
// 比如 1, 5, 3, 6，买卖两次的利润大于最大价差策略的买卖一次
// 把曲线中的每一个高点（高于相邻的两侧），视为一个卖点
// 而买点是卖点左侧最近的低点（低于相邻两侧）
// 时间复杂度O(n)，空间复杂度O(1)
func maxProfit(prices []int) (profit int) {
	n := len(prices)
	if n == 1 {
		return 0
	}
	buyDay := 0
	for buyDay < n {
		// dertermine a buy day
		for buyDay < n-1 && prices[buyDay] >= prices[buyDay+1] {
			buyDay++
		}
		if buyDay == n-1 {
			break
		}
		sellDay := buyDay + 1 // a certain sell day
		// try to find a higher price day
		for sellDay < n-1 && prices[sellDay] < prices[sellDay+1] {
			sellDay++
		}
		profit += prices[sellDay] - prices[buyDay]
		buyDay = sellDay + 1
	}
	return profit
}

// 对前一个方法进行优化，只遍历一次
// 优化的关键在于：低点-高点相当于一段递增的线段，那么我们只需要计算相邻的递增价格的收益
func maxProfitOptimization(prices []int) (profit int) {
	n := len(prices)
	if n == 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

// 跳跃游戏
// 贪心，O(n)
// 从第一个格子逐渐更新可以跳到的最远处
// 直至 到达终点 或 尝试完毕，发现终点不可达
func canJump(nums []int) (result bool) {
	n := len(nums)
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	furthest := nums[0]
	for i := 0; i <= furthest; i++ {
		furthest = getMax(furthest, i+nums[i])
		if furthest >= n-1 {
			return true
		}
	}
	return false
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// x 的平方根
func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	y := x >> 1
	for {
		tmp := y * y
		if (tmp == x) || (tmp < x && (y+1)*(y+1) > x) {
			break
		} else {
			y = (y + x/y) >> 1
		}
	}
	return y
}

// 有效的完全平方数
// 法一：二分法
func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left, right := 2, num
	for left <= right {
		mid := left + ((right - left) >> 1)
		product := mid * mid
		if product == num {
			return true
		} else if product < num {
			if (mid+1)*(mid+1) > num {
				return false
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// 法二：牛顿迭代法
// 收敛的O(logn)，优于二分法，空间复杂度更低
// best
func isPerfectSquare2(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	x := num >> 1
	for {
		product := x * x
		if product == num {
			return true
		} else if product < num && (x+1)*(x+1) > num {
			return false
		}
		x = (x + num/x) >> 1
	}

	return false
}
