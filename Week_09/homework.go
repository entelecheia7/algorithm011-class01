package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(validPalindrome2("abca"))
}

/* 简单 */
// 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// 反转字符串 II （亚马逊在半年内面试中考过）
func reverseStr(s string, k int) string {
	n := len(s)
	result := make([]byte, 0, n)
	start := 0
	reverseFlg := true
	for start < n {
		right := getMin(n-1, start+k-1)
		if reverseFlg {
			for i := right; i >= start; i-- {
				result = append(result, s[i])
			}
		} else {
			result = append(result, []byte(s[start:right+1])...)
		}
		reverseFlg = !reverseFlg
		start += k
	}

	return string(result)
}

// 翻转字符串里的单词（微软、字节跳动、苹果在半年内面试中考过）
func reverseWords1(s string) string {
	if s == "" {
		return ""
	}
	s = strings.Trim(s, " ")
	result := make([]byte, 0, len(s))
	right := len(s) // 单词的右边界，不含
	for i := len(s) - 1; i >= 0; {
		if s[i] == ' ' {
			result = append(result, []byte(s[i+1:right])...)
			result = append(result, ' ')
			i--
			for s[i] == ' ' {
				i--
			}
			right = i + 1
		} else {
			i--
		}
	}
	result = append(result, []byte(s[:right])...)
	return string(result)
}

// 反转字符串中的单词 III （微软、字节跳动、华为在半年内面试中考过）
func reverseWords2(s string) string {
	n := len(s)
	s += " "
	begin := 0 // 单词的起始位置
	result := make([]byte, 0, n)
	for i := 0; i <= n; i++ {
		if s[i] == ' ' {
			for c := i - 1; c >= begin; c-- {
				result = append(result, s[c])
			}
			result = append(result, ' ')
			begin = i + 1
		}
	}
	return string(result[:n])
}

// 仅仅反转字母（字节跳动在半年内面试中考过）
func reverseOnlyLetters(S string) string {
	length := len(S)
	if length <= 1 {
		return S
	}
	bytes := []byte(S)
	left, right := 0, length-1
	for left < right {
		if !isLetter(bytes[left]) {
			left++
		} else if !isLetter(bytes[right]) {
			right--
		} else {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}

func isLetter(x byte) bool {
	return (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')
}

// 同构字符串（谷歌、亚马逊、微软在半年内面试中考过）
// 法一：暴力，保存s=>t、t=>s的映射关系，逐一验证
func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte) // s[i] => t[i]的映射关系
	m2 := make(map[byte]byte) // t[i] => s[i]的映射关系
	for i := 0; i < len(s); i++ {
		v, exist := m1[s[i]]
		if !exist {
			m1[s[i]] = t[i]
		} else if v != t[i] {
			return false
		}
		v, exist = m2[t[i]]
		if !exist {
			m2[t[i]] = s[i]
		} else if v != s[i] {
			return false
		}

	}
	return true
}

// 法二：对于s[i]和t[i]，它们在s和t中具有相同的下标位置。
func isIsomorphic2(s string, t string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if strings.Index(s, s[i:i+1]) != strings.Index(t, t[i:i+1]) {
			return false
		}
	}
	return true
}

// 法三：法二的另一种写法
func isIsomorphic3(s string, t string) bool {
	n := len(s)
	var m1, m2 [256]int
	for i := 0; i < n; i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

// 验证回文字符串 Ⅱ（Facebook 在半年内面试中常考）
// 法一：递归
func validPalindrome(s string) bool {
	return validPalindromeHelper(s, false)
}
func validPalindromeHelper(s string, deleted bool) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			if deleted {
				return false
			}
			// 删除s[left]或s[right]
			return validPalindromeHelper(s[left+1:right+1], true) || validPalindromeHelper(s[left:right], true)
		}
		left++
		right--
	}
	return true
}

// 法二：循环
// best
func validPalindrome2(s string) (result bool) {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		// 删除s[left]
		l, r := left+1, right
		result = true
		for l < r {
			if s[l] != s[r] {
				result = false
				break
			}
			l++
			r--
		}
		if result {
			return
		}
		// 删除s[right]
		result = true
		l, r = left, right-1
		for l < r {
			if s[l] != s[r] {
				result = false
				break
			}
			l++
			r--
		}
		return
	}
	return true
}

/* 中等 */
// 最长上升子序列（字节跳动、亚马逊、微软在半年内面试中考过）
// 法一：动态规划
// dp[i]代表nums[i]结尾的最长上升子序列长度
// dp[i]的取决于dp[j]，0<=j<i 且 nums[j] < nums[i]，则dp[i] = dp[j]+1
// 时间 O(n^2)，空间 O(n)
func lengthOfLIS(nums []int) (result int) {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		result = getMax(result, dp[i])
	}
	return result
}

// 法二：dp+二分查找
// tail[i]代表长度为i+1的上升子序列中的末位最小值
// O(nlogn)
func lengthOfLIS2(nums []int) (result int) {
	n := len(nums)
	if n <= 1 {
		return n
	}
	tail := make([]int, n)
	tail[0] = nums[0]
	key := 0 // 目前计算完毕的tail[i]的索引
	for i := 1; i < n; i++ {
		if nums[i] > tail[key] {
			key++
			tail[key] = nums[i]
		} else {
			// 在计算完毕的tail范围内查找大于等于nums[i]的最小元素
			// 试图缩小计算完毕的长度为m的子序列的范围
			left, right := 0, key
			for left < right {
				mid := left + ((right - left) >> 1)
				if tail[mid] >= nums[i] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			tail[left] = nums[i]
		}
	}
	return key + 1
}

// 解码方法（字节跳动、亚马逊、Facebook 在半年内面试中考过）
// 法一：动态规划
// dp[n]表示s的前n个字符有多少种解码方式
// 如果s[i] = 0, dp[i] = dp[i-2]
// 如果s[i-1]是1，dp[i] = dp[i-2] + dp[i-1]
// 如果s[i-1]是2，且s[i]为1-6，dp[i] = dp[i-2] + dp[i-1]
// 其他情况，dp[i] = dp[i-1]
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1 // 为了方便计算
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if s[i-1] == '0' {
			// 出现'0'有两种情况，一种是10或20，一种是当前的'0'不合法
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			}
			dp[i] = dp[i-2]
		} else if (s[i-2] == '1') || (s[i-2] == '2' && s[i-1] >= '1' && s[i-1] <= '6') {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}

// 法二：对法一的空间优化，best
func numDecodings2(s string) (cur int) {
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

// 字符串转换整数 (atoi) （亚马逊、微软、Facebook 在半年内面试中考过）
func myAtoi(s string) (num int) {
	// 去掉左侧空格
	for s != "" && s[0] == ' ' {
		s = s[1:]
	}
	if s == "" {
		return 0
	}
	signed := 1 // 记录符号，去掉符号位
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			signed = -1
		}
		s = s[1:]
	}
	// 禁止出现两个符号位或非数字字符
	if s == "" || s[0] == '+' || s[0] == '-' || s[0] < '0' || s[0] > '9' {
		return 0
	}
	for s != "" {
		if s[0] < '0' || s[0] > '9' {
			break
		}
		n := int(s[0] - '0')
		if num > (math.MaxInt32-n)/10 {
			if signed == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		num = num*10 + n
		s = s[1:]
	}

	return num * signed
}

// 找到字符串中所有字母异位词（Facebook 在半年内面试中常考）
// 法一：滑动窗口
func findAnagrams(s string, p string) (result []int) {
	sl, pl := len(s), len(p)
	if pl > sl {
		return nil
	}
	need := [26]int{}
	for i := 0; i < pl; i++ {
		need[p[i]-'a']++
	}
	validType := 0
	for i := 0; i < 26; i++ {
		if need[i] > 0 {
			validType++
		}
	}
	window := [26]int{}
	valid := 0
	for i := 0; i < sl; i++ {
		// 更新窗口右侧新边界字母
		w := s[i] - 'a'
		if need[w] > 0 {
			window[w]++
			if window[w] == need[w] {
				valid++
			}
		}
		// 收缩窗口
		if i >= pl-1 {
			if valid == validType {
				result = append(result, i+1-pl)
			}
			left := s[i+1-pl] - 'a'
			if need[left] > 0 {
				if window[left] == need[left] {
					valid--
				}
				window[left]--
			}
		}
	}
	return result
}

// 最长回文子串（亚马逊、字节跳动、华为在半年内面试中常考）
// 法一：暴力+双指针
// best
func longestPalindrome(s string) (result string) {
	n := len(s)
	if n == 0 {
		return
	}
	result = s[:1]
	for i := 1; i < n-len(result)/2; i++ {
		// 以s[i]为奇数长度回文串的中心
		left, right := i-1, i+1
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if left < i-1 && right-left-1 > len(result) {
			result = s[left+1 : right]
		}
		// 以s[i]为偶数长度回文串的右中心
		left, right = i-1, i
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if right-left > 1 && right-left-1 > len(result) {
			result = s[left+1 : right]
		}
	}
	return result
}

// 法二：动态规划
// dp[i][j]表示s[i:j]是否为回文字符串，含两侧边界
// O(n^2)
// 该解法还可以进行优化，dp[i][j]仅和sp[i+1][j-1]相关
func longestPalindrome2(s string) (result string) {
	n := len(s)
	if n <= 1 {
		return s
	}
	result = s[:1]
	dp := make([][]bool, n)
	for k := range dp {
		dp[k] = make([]bool, n)
		dp[k][k] = true
	}
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j == i+1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
				if dp[i][j] && j-i+1 > len(result) {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}

/* 困难 */
// 最长有效括号（亚马逊、字节跳动、华为在半年内面试中考过）
func longestValidParentheses(s string) (maxLen int) {
	n := len(s)
	if n < 2 {
		return 0
	}
	stack := []int{-1}
	i := 0
	for i < n {
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
		i++
	}
	return maxLen
}

// 赛车（谷歌在半年内面试中考过）
// dp[i]表示target是i时的最短指令长度
// target = 2^n-1时，指令均为A，这时一定是最短指令
// 当target是其他位置时，假设2^(k-1) <= target < 2^k
// 那么可以先走(k-1)次A，到达位置(2^(k-1)-1)，通过R转向，再走m个A，再走R转向，递归走剩余路程
// 也可以走k次A，到达(2^k-1)，通过R转向，再走剩余路程:(i<<k)-1-target
var dp [10001]int

func racecar(target int) int {
	if dp[target] > 0 {
		return dp[target]
	}
	k := int(math.Floor(math.Log2(float64(target)))) + 1
	if target+1 == (1 << k) {
		return k
	}
	// 走k次A，到达(2^k-1)，通过R转向，再走剩余路程
	dp[target] = k + 1 + racecar((1<<k)-1-target)
	// 走(k-1)次A，到达位置(2^(k-1)-1)，通过R转向，再走m个A，再走R转向，递归走剩余路程
	// m的取值范围是[0, k-1)
	for m := 0; m < k-1; m++ {
		dp[target] = getMin(dp[target], k+m+1+racecar(target-(1<<(k-1))+(1<<m)))
	}
	return dp[target]
}

// 通配符匹配（Facebook、微软、字节跳动在半年内面试中考过）
// dp[i][j]表示s的前i个字符是否可以匹配p的前j个字符
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	// init
	for k := range dp {
		dp[k] = make([]bool, n+1)
	}
	// 当s为空时
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ { // p 不为空，所以从1开始
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' { // *可以匹配空字符串或者任意多个字符
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}

// 不同的子序列（MathWorks 在半年内面试中考过）
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ { // 空字符串可以视为任何字符串的子序列
		dp[0][j] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

/* helper */
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
