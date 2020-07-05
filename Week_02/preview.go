package main

import (
	"fmt"
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
// 子集
// N 皇后
