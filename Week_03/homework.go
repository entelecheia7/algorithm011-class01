package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(permute([]int{1, 2, 3}))
	// fmt.Println(combine(4, 2))

	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique2([]int{1, 1, 2}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 终结条件
	if root == nil {
		return nil
	} else if root == p || root == q {
		return root
	}
	// 处理当前层
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// 如果left和right均不为空，说明p、q分别在root的两侧子树
	return root
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	inorderMap := make(map[int]int, n)
	for i := 0; i < n; i++ {
		inorderMap[inorder[i]] = i
	}
	rootIndex := 0
	return buildTreeHelper(preorder, &rootIndex, inorder, 0, n-1, inorderMap)
}
func buildTreeHelper(preorder []int, rootIndex *int, inorder []int, left, right int, inorderMap map[int]int) *TreeNode {
	if left > right {
		return nil
	}
	root := &TreeNode{
		Val: preorder[*rootIndex],
	}
	(*rootIndex)++
	root.Left = buildTreeHelper(preorder, rootIndex, inorder, left, inorderMap[root.Val]-1, inorderMap)
	root.Right = buildTreeHelper(preorder, rootIndex, inorder, inorderMap[root.Val]+1, right, inorderMap)
	return root
}

// https://leetcode-cn.com/problems/combinations/
// 77. 组合
func combine(n int, k int) (result [][]int) {
	if n < 1 || k < 1 || n < k {
		return nil
	}
	cur := make([]int, k, k)
	combineHelper(cur, 0, 1, n, k, &result)
	return result
}

// ci 是在本次函数中，cur中要添加的元素的位置，避免append操作
func combineHelper(cur []int, ci, start, n, k int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for i := start; i <= n-k+1; i++ {
		cur[ci] = i
		combineHelper(cur, ci+1, i+1, n, k-1, result)
	}
}

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
// 47. 全排列 II
// 法一：常规回溯
func permuteUnique(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	cur := make([]int, len(nums))
	used := make([]bool, len(nums))
	permuteUniqueHelper(nums, used, cur, 0, &result)
	return result
}
func permuteUniqueHelper(nums []int, used []bool, cur []int, i int, result *[][]int) {
	if i == len(nums) {
		tmp := make([]int, i)
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	for j := 0; j < len(nums); j++ {
		if used[j] {
			continue
		}
		if j > 0 && nums[j] == nums[j-1] && !used[j-1] {
			continue
		}
		cur[i] = nums[j]
		used[j] = true
		permuteUniqueHelper(nums, used, cur, i+1, result)
		used[j] = false
	}
}

// 法二：原地算法的回溯，交换nums中的元素位置生成新的排列
func permuteUnique2(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	permuteUniqueHelper2(nums, 0, &result)
	return result
}
func permuteUniqueHelper2(nums []int, i int, result *[][]int) {
	if i == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	// nums[i]和其他不同元素交换位置
	for j := i; j < len(nums); j++ {
		if j != i && nums[j] == nums[i] { // 跳过重复元素
			continue
		}
		nums[j], nums[i] = nums[i], nums[j]
		permuteUniqueHelper2(nums, i+1, result)
	}
	// 还原
	for k := len(nums) - 1; k > i; k-- {
		nums[k], nums[i] = nums[i], nums[k]
	}
}

// https://leetcode-cn.com/problems/majority-element/description/

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// https://leetcode-cn.com/problems/n-queens/
