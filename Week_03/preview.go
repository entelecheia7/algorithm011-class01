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
// 买卖股票的最佳时机 II
// 跳跃游戏
// x 的平方根
// 有效的完全平方数
