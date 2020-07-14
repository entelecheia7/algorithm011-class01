package main

// import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的层次遍历
func levelOrder(root *TreeNode) (result [][]int) {
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
// x 的平方根
// 有效的完全平方数
