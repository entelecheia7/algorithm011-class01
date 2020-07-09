package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(climbStairs(5))

	fmt.Println(generateParenthesis2(3))
}

// 爬楼梯
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	pp := 1 // f(n-2)
	p := 2  // f(n-1)
	for i := 3; i <= n; i++ {
		p, pp = p+pp, p
	}
	return p
}

// 括号生成
// 法一：回溯
// 法二：动态规划
func generateParenthesis1(n int) (result []string) {
	if n <= 0 {
		return nil
	}
	generateParenthesisHelper("(", &result, n-1, n)
	return result
}
func generateParenthesisHelper(cur string, result *[]string, left, right int) {
	if left == 0 && right == 0 {
		*result = append(*result, cur)
		return
	}
	if left > 0 {
		generateParenthesisHelper(cur+"(", result, left-1, right)
	}
	if right > left {
		generateParenthesisHelper(cur+")", result, left, right-1)
	}
}
func generateParenthesis2(n int) []string {
	if n <= 0 {
		return nil
	}
	state := make([][]string, n+1)
	state[0] = []string{""}
	for i := 1; i <= n; i++ {
		cur := []string{}
		for j := 0; j < i; j++ {
			subState1 := state[j]
			subState2 := state[i-j-1]
			for _, str1 := range subState1 {
				for _, str2 := range subState2 {
					cur = append(cur, "("+str1+")"+str2)
				}
			}
		}
		state[i] = cur
	}
	return state[n]
}

// Pow(x, n)
// 法一：递归，分治，O(logn)
// 注意考虑-n的溢出情况
func myPow1(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n > 0 {
		return myPowHelper(x, n)
	}
	if n == math.MinInt32 {
		x *= x
		n >>= 1
	}
	return 1.0 / myPowHelper(x, -n)
}
func myPowHelper(x float64, n int) (result float64) {
	if n == 0 || x == 1 {
		return 1
	}
	half := myPowHelper(x, n>>1)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// 法二：循环，二进制解法
// O(logn)
func myPow2(x float64, n int) (result float64) {
	if n == 0 || x == 1 {
		return 1
	}
	if n == math.MinInt32 {
		x *= x
		n >>= 1
	}
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	result = 1
	factor := x
	for n > 0 {
		if n%2 == 1 {
			result = result * factor
		}
		factor *= factor
		n >>= 1
	}
	return result
}

// 子集
// 法一：回溯，逐个生成长度为1、2、3……len(nums)的子集
func subsets1(nums []int) (result [][]int) {
	subsetBTHelper(nums, 0, []int{}, &result)

	return result
}
func subsetBTHelper(nums []int, from int, cur []int, result *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*result = append(*result, tmp)

	for i := from; i < len(nums); i++ {
		subsetBTHelper(nums, i+1, append(cur, nums[i]), result)
	}
}

// 法二：二进制
// 将 nums 视为长度为 len(nums) 的二进制串 x，则 nums 的子集就是 0-n 位的 x 掩码
// 以 length 为 3 的 nums 为例，其二进制位掩码分别为 000、001、010、011、100、101、111。
// 通过计算其掩码的位是否为 1 (i >> j & 1)来决定是否添加 nums 对于位置的数字到 tmp，得出子集。
func subsets2(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	for i := 0; i < (1 << len(nums)); i++ { // 本层循环用于提供掩码
		tmp := []int{}
		for j := 0; j < len(nums); j++ {
			if (i >> j & 1) == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		result = append(result, tmp)
	}

	return result
}

// N 皇后
func solveNQueens(n int) (result [][]string) {
	// 生成空棋盘
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	nQueensHelper(n, 0, board, &result)
	return result
}

// row代表放置的行，范围：0~n-1
func nQueensHelper(n, row int, board []string, result *[][]string) {
	if row == n {
		tmp := make([]string, n)
		copy(tmp, board)
		*result = append(*result, tmp)
		return
	}
	for column := 0; column < n; column++ {
		if checkNQueen(n, row, column, board) {
			initRow := board[row]
			rowStr := []byte(board[row])
			rowStr[column] = 'Q'
			board[row] = string(rowStr)
			nQueensHelper(n, row+1, board, result)
			board[row] = initRow
		}
	}

}

// row和column代表新皇后想要放置的行和列
func checkNQueen(n, row, column int, cur []string) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- { // i表示已放置皇后的行
		if cur[i][column] == 'Q' { // 检查竖行
			return false
		}
		if leftup >= 0 && cur[i][leftup] == 'Q' { // 检查左斜线
			return false
		}
		if rightup < n && cur[i][rightup] == 'Q' { // 检查右斜线
			return false
		}
		leftup--
		rightup++
	}
	return true
}
