package main

import (
// "fmt"
)

func main() {
}

/* 简单 */

// 柠檬水找零（亚马逊在半年内面试中考过）
// 找零时，总是试图先支出最大的面值
func lemonadeChange(bills []int) bool {
	n := len(bills)
	if n == 0 {
		return true
	} else if bills[0] != 5 {
		return false
	}
	count5, count10 := 1, 0
	for i := 1; i < n; i++ {
		if bills[i] == 5 { // 不找零
			count5++
		} else if bills[i] == 10 { // 找5块
			if count5 == 0 {
				return false
			}
			count10++
			count5--
		} else { // 找15
			if count10 > 0 && count5 > 0 {
				count10--
				count5--
			} else if count5 >= 3 {
				count5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// 买卖股票的最佳时机 II （亚马逊、字节跳动、微软在半年内面试中考过）
// 分发饼干（亚马逊在半年内面试中考过）
// 模拟行走机器人
// 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
// 说明：同学们可以将自己的思路、代码写在第 4 周的学习总结中

/* 中等 */

// 单词接龙（亚马逊在半年内面试常考）
// 岛屿数量（近半年内，亚马逊在面试中考查此题达到 350 次）
// 扫雷游戏（亚马逊、Facebook 在半年内面试中考过）
// 跳跃游戏 （亚马逊、华为、Facebook 在半年内面试中考过）
// 搜索旋转排序数组（Facebook、字节跳动、亚马逊在半年内面试常考）
// 搜索二维矩阵（亚马逊、微软、Facebook 在半年内面试中考过）
// 寻找旋转排序数组中的最小值（亚马逊、微软、字节跳动在半年内面试中考过）

/* 困难 */

// 单词接龙 II （微软、亚马逊、Facebook 在半年内面试中考过）
// 跳跃游戏 II （亚马逊、华为、字节跳动在半年内面试中考过）
