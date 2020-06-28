package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
}

/* 简单 */
// 删除排序数组中的重复项（Facebook、字节跳动、微软在半年内面试中考过）
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

// 旋转数组（微软、亚马逊、PayPal 在半年内面试中考过）
// 法一：借助一个额外的数组，空间复杂度O(n)，时间复杂度O(n)，非原地算法，略
// 法二：每次右移1位，重复k次，空间复杂度O(1)，时间复杂度O(k*n)，原地算法
func rotate2(nums []int, k int) {
	n := len(nums)
	if n < 2 {
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

// 法三：反转数组。实际上
func rotate3(nums []int, k int) {
	n := len(nums)
	if n < 2 {
		return
	}

}

// 合并两个有序链表（亚马逊、字节跳动在半年内面试常考）
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

// 合并两个有序数组（Facebook 在半年内面试常考）
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

// 两数之和（亚马逊、字节跳动、谷歌、Facebook、苹果、微软在半年内面试中高频常考）

// 移动零（Facebook、亚马逊、苹果在半年内面试中考过）

// 加一（谷歌、字节跳动、Facebook 在半年内面试中考过）

/* 中等 */
// 设计循环双端队列（Facebook 在 1 年内面试中考过）

/* 困难 */
// 接雨水（亚马逊、字节跳动、高盛集团、Facebook 在半年内面试常考）

type ListNode struct {
	Val  int
	Next *ListNode
}
