package main

func main() {

}

// 1.有效的字母异位词
func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < ls; i++ {
		if s[i] != t[i] {
			m[s[i]-'a']++
			m[t[i]-'a']--
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// 法二：给两个字符串按同一算法排序，看生成的字符串是否相等
// 进阶：将数组换成map，map的下标可以是unicode字符

// 2.二叉树的中序遍历

// 3.最小的 k 个数
