package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
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

	// 矩形区域不超过 K 的最大数值和

}

/* 中等 */

// 最小路径和（亚马逊、高盛集团、谷歌在半年内面试中考过）
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = grid[i][j] + dp[j]
			} else {
				dp[j] = getMin(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

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
// 法一：暴力，best
// 以s中的每个字符为回文串中点，检查计算
func countSubstrings1(s string) (count int) {
	n := len(s)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		count++
		// 回文串长度为奇数
		count += extendSubString(s, n, i-1, i+1)
		// 回文串长度为偶数，以s[i]为中点左侧字符
		count += extendSubString(s, n, i, i+1)
	}
	return count
}
func extendSubString(s string, n, left, right int) (count int) {
	for left >= 0 && right < n && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

// 法二：动态规划
// 长度更长的回文串总是在长度稍短的回文串的基础上形成
// dp[i][j]表示 s[i:j]是否为回文子串
// dp[i][j]在 dp[i+1][j-1]的基础上判断扩展
func countSubstrings2(s string) (count int) {
	n := len(s)
	if n <= 1 {
		return
	}
	dp := make([][]bool, n)
	for k := range dp {
		dp[k] = make([]bool, n)
	}
	for j := 0; j < n; j++ { // j包裹i循环是为了保证 dp[i+1][j-1] 已经算出
		for i := j; i >= 0; i-- {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}

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
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	} else if l2 == 0 {
		return l1
	}
	// init
	// 初始化 word1[:0……l1]到word2[:0]的编辑距离
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	// 初始化 word1[:0]到word2[:0……l2]的编辑距离
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	// 开始递推
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin(getMin(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}

	return dp[l1][l2]
}

// 矩形区域不超过 K 的最大数值和（谷歌在半年内面试中考过）

// 青蛙过河（亚马逊、苹果、字节跳动在半年内面试中考过）
// 法一：动态规划
// dp[i][j]表示在从某个位置 x 跳 j 步 是否可达stones[i]（1<=j<=i）
// 则到达位置 x 的步数为 j-1 || j || j+1
// dp[i][j] = dp[x][k-1] || dp[x][j] || dp[x][j+1]
// O(n^2)
func canCross(stones []int) bool {
	n := len(stones)
	if n == 2 {
		if stones[1] != 1 {
			return false
		}
		return true
	}
	dp := make([][]bool, n)
	for k := range dp {
		dp[k] = make([]bool, n+1) // +1是为了保证 j+1 步不越界
	}
	dp[1][1] = true
	for i := 2; i < n; i++ {
		// 本次循环求解 dp[i]，stones[i]可以从 1~i-1 位置抵达，列举所有可能
		for x := 1; x < i; x++ {
			// 从 stones[x] 跳到 stones[i]需要多少步
			needStep := stones[i] - stones[x]
			// 如果从 stones[x] 跳 needStep 步可达 stones[i]，则有 dp[x][needStep] 或 dp[x][needStep-1] 或 dp[x][needStep+1] 为 true
			if needStep <= i {
				dp[i][needStep] = dp[x][needStep] || dp[x][needStep-1] || dp[x][needStep+1]
				if i == n-1 && dp[i][needStep] {
					return true
				}
			}
		}

	}

	return false
}

// 分割数组的最大值（谷歌、亚马逊、Facebook 在半年内面试中考过）
// 法一：动态规划，O(n^3)
// dp[i][j]表示以将nums的前i个数分为 j 组得到的最大连续子数组和的最小值(j<=i)
// 设前 k 个数分为 j-1 组，最后一组为第 k+1 到第 i 个数 (k+1>=j)
func splitArray1(nums []int, m int) (result int) {
	n := len(nums)
	dp := make([][]int, n+1)
	// init
	for k := range dp {
		dp[k] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[k][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0
	// sums 表示前i个元素的和
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m && j <= i; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = getMin(dp[i][j], getMax(sums[i]-sums[k], dp[k][j-1]))
			}
		}
	}

	return dp[n][m]
}

// 法二：二分查找
// best
// nums子数组的值在 [max(nums), sum(nums)]中间，使用二分查找不断查找接近值
func splitArray2(nums []int, m int) (result int) {
	n := len(nums)
	left := nums[0]
	right := nums[0]
	for i := 1; i < n; i++ {
		right += nums[i]
		if nums[i] > left {
			left = nums[i]
		}
	}
	for left < right {
		mid := left + ((right - left) >> 1)
		// 从nums[0]开始计算一个 subSum
		// 每找出一个subSum <= mid 的子数组，统计数量+1
		count := 1 // 循环中subSum超过mid才进行统计，因此最后一轮没有在循环中统计，所以赋1
		subSum := 0
		for i := 0; i < n; i++ {
			subSum += nums[i]
			if subSum > mid {
				count++
				subSum = nums[i]
			}
		}
		if count > m { // 分组数量过多，说明选择的subSum过小
			left = mid + 1
		} else { // 分组数量小于等于m，说明选择的subSum可能过大
			right = mid
		}
	}

	return left
}

// 学生出勤记录 II （谷歌在半年内面试中考过）
// 法一：回溯，时间复杂度高，无法AC
// 法二：动态规划
// 将可能性分为6种：
// a 不含A的LL结尾，可添加A P
// b 不含A的L结尾，倒数第二位不是L，可添加A L P
// c 不含A的非L结尾，可添加A L P
// d 含A的LL结尾，可添加P
// e 含A的L结尾，倒数第二位不为L，可添加 L P
// f 含A的非L结尾，可添加 L P
func checkRecord(n int) int {
	// 初始化n=1
	a, b, c, d, e, f := 0, 1, 1, 0, 0, 1
	// i 本该从 2 开始
	// 但由于最后的结果是(a + b + c + d + e + f) % 1000000007
	// f恰好是a-f的和，所以直接多循环一遍
	for i := 1; i <= n; i++ {
		a, b, c = b, c, (a+b+c)%1000000007
		d, e, f = e, f, (d+e+f+c)%1000000007
	}

	return f
}

// 最小覆盖子串（Facebook 在半年内面试中常考）
// 法一：暴力搜索，时间复杂度高，O(m*n)
// 法二：滑动窗口
// 先确定窗口大小：固定左侧，右侧滑动，当s包含t的所有元素时，就确定了初始窗口大小
// 当窗口大小确定时，开始记录窗口大小并试图移动左侧窗口
func minWindow(s string, t string) (result string) {
	ls, lt := len(s), len(t)
	if ls < lt || lt == 0 || ls == 0 {
		return
	}
	var need, window [128]int // map是一种更通用的做法，但数组的效率更高
	for i := 0; i < lt; i++ {
		need[t[i]]++ // 记录需要的元素及数量
	}
	needTypeLen := 0 // 记录需要的不同元素的数量
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			needTypeLen++
		}
	}
	valid := 0
	left, right := 0, 0
	for right < ls {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 窗口包含所有所需元素时，进行记录，并移动左侧边界
		for valid == needTypeLen {
			// 计算长度
			if result == "" || right-left < len(result) {
				result = s[left:right]
			}
			// 滑动左侧边界
			c = s[left]
			if need[c] > 0 {
				if need[c] == window[c] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}

	return
}

// 戳气球（亚马逊在半年内面试中考过）
// 法一：回溯，时间复杂度高，略
// 法二：动态规划
// dp[i][j] 表示，戳破气球 i 和气球 j 之间（不含 i、 j）的所有气球，可获得的最高分数。
// 设 nums (i, j)间最后一个被戳的气球是 k，则 dp[i][j] = dp[i][k] + nums[k]*nums[i]*nums[j]+dp[k][j]
// 固定 i，j，对 k 进行枚举
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	balls := make([]int, n+2)
	balls[0] = 1
	copy(balls[1:], nums)
	balls[n+1] = 1

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	// 对于任意的 dp[i][j]，需要满足 dp[i][k] 和 dp[k][j]已被计算
	// i < k < j
	// 因此 i 需要倒序计算，j从左到右计算
	for i := n + 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = getMax(dp[i][k]+balls[k]*balls[i]*balls[j]+dp[k][j], dp[i][j])
			}
		}
	}

	return dp[0][n+1]
}

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
