package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// fmt.Println(permute([]int{1, 2, 3}))
	// fmt.Println(combine(4, 2))

	// fmt.Println(permuteUnique([]int{1, 1, 2}))
	// fmt.Println(permuteUnique2([]int{1, 1, 2}))

	fmt.Println(solveNQueens(4))
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
// 169. 多数元素
// 法一：借助一个map统计元素频率，O(n)
func majorityElement1(nums []int) int {
	freq := len(nums) / 2
	m := make(map[int]int, freq)
	for _, num := range nums {
		m[num]++
		if m[num] > freq {
			return num
		}
	}
	return -1
}

// 法二：Boyer-Moore 投票算法
// 时间复杂度O(n)，空间复杂度O(1)
func majorityElement2(nums []int) int {
	candidate, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count++
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合
var digitsRelation = [][]string{
	0: nil,
	1: nil,
	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	6: []string{"m", "n", "o"},
	7: []string{"p", "q", "r", "s"},
	8: []string{"t", "u", "v"},
	9: []string{"w", "x", "y", "z"},
}

// 法一：循环版，空间复杂度大
func letterCombinations1(digits string) (result []string) {
	n := len(digits)
	if n == 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			result = digitsRelation[digits[i]-'0']
		} else {
			tmp := make([]string, 0, len(result)*len(digitsRelation[digits[i]-'0']))
			for _, r := range result {
				for _, s := range digitsRelation[digits[i]-'0'] {
					tmp = append(tmp, r+s)
				}
			}
			result = tmp
		}
	}
	return result
}

// 法二：回溯
// best
func letterCombinations2(digits string) (result []string) {
	n := len(digits)
	if n == 0 {
		return nil
	}
	letterCombinationsHelper(digits, "", &result)
	return result
}

func letterCombinationsHelper(digits string, cur string, result *[]string) {
	if digits == "" {
		*result = append(*result, cur)
		return
	}
	for _, s := range digitsRelation[digits[0]-'0'] {
		letterCombinationsHelper(digits[1:], cur+s, result)
	}
}

// https://leetcode-cn.com/problems/n-queens/
// 51. N皇后
func solveNQueens(n int) (result [][]string) {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	solveNQueensHelper(n, board, 0, &result)
	return result
}
func solveNQueensHelper(n int, board []string, row int, result *[][]string) {
	if row == n {
		tmp := make([]string, n)
		copy(tmp, board)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if checkNQueenPos(board, row, i) {
			orig := board[row]
			putQueen := []byte(board[row])
			putQueen[i] = 'Q'
			board[row] = string(putQueen)
			solveNQueensHelper(n, board, row+1, result)
			board[row] = orig
		}
	}
}

// 检查当前想要放置Queen的位置是否合法
func checkNQueenPos(board []string, row, column int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		if board[i][column] == 'Q' {
			return false
		}
		if leftup >= 0 && board[i][leftup] == 'Q' {
			return false
		}
		if rightup < len(board) && board[i][rightup] == 'Q' {
			return false
		}
		leftup--
		rightup++

	}
	return true
}
