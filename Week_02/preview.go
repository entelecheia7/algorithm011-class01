package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(5))
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

// Pow(x, n)
// 子集
// N 皇后
