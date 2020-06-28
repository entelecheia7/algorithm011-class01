package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate4(nums, 3)
	fmt.Println(nums)
}

/* 简单 */
// 1.删除排序数组中的重复项（Facebook、字节跳动、微软在半年内面试中考过）
func removeDuplicates(nums []int) int {
	n := len(nums)
	j := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// 2.旋转数组（微软、亚马逊、PayPal 在半年内面试中考过）
// 法一：借助一个额外的数组，空间复杂度O(n)，时间复杂度O(n)，非原地算法，略
// 法二：每次右移1位，重复k次，空间复杂度O(1)，时间复杂度O(k*n)，原地算法
func rotate2(nums []int, k int) {
	n := len(nums)
	if n < 2 || k == 0 {
		return
	}
	tmp := nums[0]
	for k > 0 {
		for i := 0; i < n; i++ {
			if i == n-1 {
				nums[0] = tmp
			} else {
				tmp, nums[i+1] = nums[i+1], tmp
			}
		}
		k--
	}
}

// 法三：三次数组反转。
// 旋转k次，数组末尾的 k%n 个元素会移动到数组头部，其余元素右移
// 首先将整个数组反转，然后反转前k个元素，再反转剩下的元素
// 空间O(1)，时间O(n)，原地算法
func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n
	if n < 2 || k == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// 法四：环状替代
func rotate4(nums []int, k int) {
	n := len(nums)
	k %= n
	if n < 2 || k == 0 {
		return
	}
	count := 0
	for i := 0; count < n; i++ {
		cur, val := i, nums[i]
	C:
		{
			next := (cur + k) % n
			tmp := nums[next]
			nums[next] = val
			cur, val = next, tmp
			count++
			if cur != i {
				goto C
			}
		}
	}
}

// 3.合并两个有序链表（亚马逊、字节跳动在半年内面试常考）
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	new := &ListNode{}
	head := new
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			new.Next = l1
			l1 = l1.Next
		} else {
			new.Next = l2
			l2 = l2.Next
		}
		new = new.Next
	}
	if l1 != nil {
		new.Next = l1
	}
	if l2 != nil {
		new.Next = l2
	}

	return head.Next
}

// 4.合并两个有序数组（Facebook 在半年内面试常考）
func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[pos] = nums1[m]
			m--
		} else {
			nums1[pos] = nums2[n]
			n--
		}
		pos--
	}
	for ; n >= 0; n-- {
		nums1[pos] = nums2[n]
		pos--
	}
}

// 5.两数之和（亚马逊、字节跳动、谷歌、Facebook、苹果、微软在半年内面试中高频常考）
func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	state := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if v, exist := state[nums[i]]; exist {
			return []int{v, i}
		}
		state[target-nums[i]] = i
	}

	return nil
}

// 6.移动零（Facebook、亚马逊、苹果在半年内面试中考过）
func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			if i != j {
				nums[j], nums[i] = nums[i], 0
			}
			j++
		}
	}
}

// 7.加一（谷歌、字节跳动、Facebook 在半年内面试中考过）
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

/* 中等 */
// 8.设计循环双端队列（Facebook 在 1 年内面试中考过）

/* 困难 */
// 9.接雨水（亚马逊、字节跳动、高盛集团、Facebook 在半年内面试常考）

type ListNode struct {
	Val  int
	Next *ListNode
}
