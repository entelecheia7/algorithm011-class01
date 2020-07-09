package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

// }

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

// https://leetcode-cn.com/problems/combinations/

// https://leetcode-cn.com/problems/permutations/
// 46. 全排列
func permute(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	used := make([]int, len(nums))
	cur := make([]int, len(nums))
	permuteHelper(nums, used, cur, 0, &result)
	return result
}
func permuteHelper(nums []int, used []int, cur []int, index int, result *[][]int) {
	if index == len(nums) {
		tmp := make([]int, index)
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		cur[index] = nums[i]
		used[i] = 1
		permuteHelper(nums, used, cur, index+1, result)
		used[i] = 0
	}
}

// https://leetcode-cn.com/problems/permutations-ii/

// https://leetcode-cn.com/problems/majority-element/description/

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// https://leetcode-cn.com/problems/n-queens/
