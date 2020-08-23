package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(validPalindrome2("abca"))

	// 在学习总结中，写出不同路径 2 这道题目的状态转移方程。
	// 最长上升子序列（字节跳动、亚马逊、微软在半年内面试中考过）
	// 解码方法（字节跳动、亚马逊、Facebook 在半年内面试中考过）
	// 字符串转换整数 (atoi) （亚马逊、微软、Facebook 在半年内面试中考过）
	// 找到字符串中所有字母异位词（Facebook 在半年内面试中常考）
	// 最长回文子串（亚马逊、字节跳动、华为在半年内面试中常考）

	// 最长有效括号（亚马逊、字节跳动、华为在半年内面试中考过）
	// 赛车（谷歌在半年内面试中考过）
	// 通配符匹配（Facebook、微软、字节跳动在半年内面试中考过）
	// 不同的子序列（MathWorks 在半年内面试中考过）

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

// 在学习总结中，写出不同路径 2 这道题目的状态转移方程。
// 最长上升子序列（字节跳动、亚马逊、微软在半年内面试中考过）
// 解码方法（字节跳动、亚马逊、Facebook 在半年内面试中考过）
// 字符串转换整数 (atoi) （亚马逊、微软、Facebook 在半年内面试中考过）
// 找到字符串中所有字母异位词（Facebook 在半年内面试中常考）
// 最长回文子串（亚马逊、字节跳动、华为在半年内面试中常考）

/* 困难 */
// 最长有效括号（亚马逊、字节跳动、华为在半年内面试中考过）
// 赛车（谷歌在半年内面试中考过）
// 通配符匹配（Facebook、微软、字节跳动在半年内面试中考过）
// 不同的子序列（MathWorks 在半年内面试中考过）

/* helper */
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
