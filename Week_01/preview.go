package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLeastNumbers1([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers2([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers3([]int{3, 2, 1}, 2))
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
// 进阶：将数组换成map，map的下标是unicode字符

// 2.二叉树的中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

// 法二：循环
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}

// 法三：莫里斯遍历。改变树的形态，将节点以中序遍历的顺序变化成只有右子节点的二叉树
func inorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Left == nil {
			// 左子节点为空，处理根节点和右子树
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			// 找到左子树的最右节点pre，将根节点和右子树添加为pre的右子树
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				// 重复地处理左子节点操作
				cur = cur.Left
			}
			if pre.Right == cur {
				// 还原树的结构
				result = append(result, cur.Val)
				pre.Right = nil
				cur = cur.Right
			}
		}
	}
	return result
}

type colorNode struct {
	node  *TreeNode
	color int
}

// 法四：颜色遍历，二叉树遍历通用
// 定义两种颜色，白色为首次访问，灰色为第二次，灰色节点可以直接输出
func inorderTraversal4(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	white, gray := 0, 1
	stack := []colorNode{{root, white}}
	result := []int{}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.node != nil {
			if cur.color == white { // 第一次访问
				stack = append(stack, colorNode{cur.node.Right, white})
				cur.color = gray
				stack = append(stack, cur)
				stack = append(stack, colorNode{cur.node.Left, white})
			} else {
				result = append(result, cur.node.Val)
			}
		}
	}
	return result
}

// 3.最小的 k 个数
// 法一：排序，时间O(nlogn)，空间O(logn)
// 法二：利用大顶堆，时间O(nlogk)，空间O(k)
// 法三：利用快排思想进行递归分区，时间O(n)
func getLeastNumbers1(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	var h maxIntHeap = make([]int, k)
	copy(h, arr)
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		if h[0] > arr[i] {
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}

	return h
}

type maxIntHeap []int

func (h maxIntHeap) Len() int            { return len(h) }
func (h maxIntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxIntHeap) Push(x interface{}) { (*h) = append(*h, x.(int)) }
func (h *maxIntHeap) Pop() interface{}   { min := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return min }

func getLeastNumbers3(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	} else if k == 0 {
		return nil
	}
	return quickSearch(arr, 0, len(arr)-1, k)
}
func quickSearch(arr []int, left, right, k int) []int {
	mid := partition(arr, left, right)
	if mid == k-1 {
		return arr[:k]
	} else if mid < k-1 {
		return quickSearch(arr, mid+1, right, k)
	}
	return quickSearch(arr, left, mid-1, k)
}

// 返回分区下标
func partition(arr []int, left, right int) (pivot int) {
	standard := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < standard {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[right], arr[i] = arr[i], arr[right]
	return i
}
