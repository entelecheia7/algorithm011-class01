package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(isAnagram1("anagram", "anagram"))
	// fmt.Println(isAnagram2("anagram", "anagram"))

	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

/* 简单： */
// 有效的字母异位词（亚马逊、Facebook、谷歌在半年内面试中考过）
// 法一：对字符串进行排序，判断相等
// 法二：统计字符出现的频率，进行比较
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
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
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	b1 := SortByte(s)
	sort.Sort(b1)
	b2 := SortByte(s)
	sort.Sort(b2)
	if b1.String() != b2.String() {
		return false
	}
	return true
}

type SortByte []byte

func (s SortByte) Less(i, j int) bool { return s[i] < s[j] }
func (s SortByte) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByte) Len() int           { return len(s) }
func (s SortByte) String() string     { return string(s) }

// 两数之和（近半年内，亚马逊考查此题达到 216 次、字节跳动 147 次、谷歌 104 次，Facebook、苹果、微软、腾讯也在近半年内面试常考）
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, num := range nums {
		if i, ok := m[num]; ok {
			return []int{i, k}
		}
		m[target-num] = k
	}
	return nil
}

// N 叉树的前序遍历（亚马逊在半年内面试中考过）
// 法一：递归
// 法二：栈+循环
type Node struct {
	Val      int
	Children []*Node
}

func preorder1(root *Node) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	for _, child := range root.Children {
		result = append(result, preorder1(child)...)
	}
	return result
}
func preorder2(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	result := []int{}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, back.Val)
		for i := len(back.Children) - 1; i >= 0; i-- {
			stack = append(stack, back.Children[i])
		}
	}

	return result
}

/* 中等： */
// 字母异位词分组（亚马逊在半年内面试中常考）
// 法一：计算每个字符串的频率
// 法二：排序
func groupAnagrams1(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for i := 0; i < len(str); i++ {
			count[str[i]-'a']++
		}
		m[count] = append(m[count], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
func groupAnagrams2(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	m := make(map[string][]string)
	for _, str := range strs {
		b := SortByte(str)
		sort.Sort(b)
		bs := b.String()
		m[bs] = append(m[bs], str)
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// 二叉树的中序遍历（亚马逊、字节跳动、微软在半年内面试中考过）
// 法一：递归
// 法二：循环
// 法三：颜色遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	result = inorderTraversal1(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal1(root.Right)...)
	return result
}
func inorderTraversal2(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return
}

type colorNode struct {
	node  *TreeNode
	color int
}

func inorderTraversal3(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	white, gray := 0, 1
	stack := []colorNode{{root, white}}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if back.node != nil {
			if back.color == white {
				stack = append(stack, colorNode{back.node.Right, white})
				back.color = gray
				stack = append(stack, back)
				stack = append(stack, colorNode{back.node.Left, white})
			} else {
				result = append(result, back.node.Val)
			}
		}
	}

	return result
}

// 二叉树的前序遍历（字节跳动、谷歌、腾讯在半年内面试中考过）
// 法一：递归
// 法二：循环
// 法三：颜色遍历
func preorderTraversal1(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	result = []int{root.Val}
	result = append(result, preorderTraversal1(root.Left)...)
	result = append(result, preorderTraversal1(root.Right)...)
	return
}
func preorderTraversal2(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, back.Val)
		if back.Right != nil {
			stack = append(stack, back.Right)
		}
		if back.Left != nil {
			stack = append(stack, back.Left)
		}
	}
	return
}
func preorderTraversal3(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	white, gray := 0, 1
	stack := []colorNode{{root, white}}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if back.node != nil {
			if back.color == white {
				stack = append(stack, colorNode{back.node.Right, white})
				stack = append(stack, colorNode{back.node.Left, white})
				back.color = gray
				stack = append(stack, back)
			} else {
				result = append(result, back.node.Val)
			}
		}
	}
	return
}

// N 叉树的层序遍历（亚马逊在半年内面试中考过）
// 法一：循环
// 法二：递归
func levelOrder1(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		n := len(queue) // 本层节点数量
		val := []int{}  // 本层节点值
		for i := 0; i < n; i++ {
			val = append(val, queue[i].Val)
			if queue[i].Children != nil {
				queue = append(queue, queue[i].Children...)
			}
		}
		queue = queue[n:]
		result = append(result, val)
	}

	return
}
func levelOrder2(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	helper(root, 0, &result)
	return
}
func helper(root *Node, level int, result *[][]int) {
	if len(*result) <= level {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	for _, child := range root.Children {
		helper(child, level+1, result)
	}
}

// 丑数（字节跳动在半年内面试中考过）
func isUgly(num int) bool {
	if num > 0 {
		for i := 2; i < 6; i++ {
			for num%i == 0 {
				num /= i
			}
		}
	}
	return num == 1
}

// 前 K 个高频元素（亚马逊在半年内面试中常考）
// 法一：利用桶排序。先使用map统计频率，再按照频率分为1-m个桶。这种方法适合频率分布范围不大且比较均匀的情况
// 法二：最小堆，统计频率后维护一个大小为k的最小堆
func topKFrequent1(nums []int, k int) (result []int) {
	frequency := make(map[int]int, k) // 数字=>频率
	for _, num := range nums {
		frequency[num]++
	}
	freq := make(map[int][]int) // 频率=>数字集合
	maxBucket := 0
	for num, f := range frequency {
		freq[f] = append(freq[f], num)
		if f > maxBucket {
			maxBucket = f
		}
	}
	for i := maxBucket; i >= 1 && k > 0; i-- {
		if num, ok := freq[i]; ok {
			if len(num) <= k {
				result = append(result, num...)
				k -= len(num)
			} else {
				result = append(result, num[:k]...)
				k = 0
			}
		}
	}
	return
}

func topKFrequent2(nums []int, k int) []int {
	frequency := make(map[int]int, k) // 数字=>频率
	for _, num := range nums {
		frequency[num]++
	}
	// 最小堆中保存数字，使用频率作为比较的依据
	h := newminIntHeap(k)
	for num, freq := range frequency {
		if h.Len() < k {
			heap.Push(h, numWithFreq{num, freq})
		} else if h.Peek().(numWithFreq).freq < freq {
			heap.Pop(h)
			heap.Push(h, numWithFreq{num, freq})
		}
	}

	result := make([]int, 0, k)
	for _, v := range h.data {
		result = append(result, v.num)
	}
	return result
}

type minIntHeap struct {
	data []numWithFreq
}
type numWithFreq struct {
	num  int
	freq int
}

func newminIntHeap(capacity int) *minIntHeap {
	return &minIntHeap{data: make([]numWithFreq, 0, capacity)}
}
func (h minIntHeap) Len() int           { return len(h.data) }
func (h minIntHeap) Less(i, j int) bool { return h.data[i].freq < h.data[j].freq }
func (h minIntHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *minIntHeap) Push(x interface{}) {
	h.data = append(h.data, x.(numWithFreq))
}
func (h *minIntHeap) Pop() interface{} {
	min := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return min
}
func (h minIntHeap) Peek() interface{} {
	if h.Len() == 0 {
		return -1
	}
	return h.data[0]
}
