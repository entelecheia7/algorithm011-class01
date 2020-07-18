package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7

	// fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1})) // 1

	fmt.Println(findDisorderIndex([]int{4, 5, 6, 7, 0, 1, 2})) // 4
	fmt.Println(findDisorderIndex([]int{4, 5, 7, 0, 1, 2}))    // 3
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
// 只有一段递增的序列才能提供利润
func maxProfit(prices []int) (profit int) {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// 分发饼干（亚马逊在半年内面试中考过）
// g代表小朋友，s代表饼干
func findContentChildren(g []int, s []int) (satisfy int) {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	si, gi := 0, 0
	for si < len(s) && gi < len(g) {
		if s[si] >= g[gi] {
			satisfy++
			gi++
		}
		si++

	}
	return satisfy
}

// 模拟行走机器人
func robotSim(commands []int, obstacles [][]int) (area int) {
	n := len(commands)
	if n == 0 {
		return 0
	}
	x, y := 0, 0
	obstaclesMap := make(map[[2]int]bool) // 将障碍物转换为map，加速查找
	for _, v := range obstacles {
		obstaclesMap[[2]int{v[0], v[1]}] = true
	}
	// 0北 1东 2南 3西
	direction := 0
	directionX := [4]int{0, 1, 0, -1} // 朝向对应的x坐标变化
	directionY := [4]int{1, 0, -1, 0} // 朝向对应的y坐标变化

	for i := 0; i < n; i++ {
		if commands[i] == -2 {
			direction = (direction + 3) % 4
		} else if commands[i] == -1 {
			direction = (direction + 1) % 4
		} else {
			for j := 0; j < commands[i]; j++ {
				tmpX, tmpY := x+directionX[direction], y+directionY[direction]
				if _, exist := obstaclesMap[[2]int{tmpX, tmpY}]; exist {
					break
				}
				x, y = tmpX, tmpY
				// 更新面积
				area = getMax(area, x*x+y*y)
			}
		}
	}
	return area
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
func findDisorderIndex(nums []int) int {
	left, right := 0, len(nums)-1
	if right <= 0 {
		return -1
	}
	for left < right {
		if right-left == 1 {
			if nums[left] > nums[right] {
				return right
			} else {
				return -1
			}
		}
		mid := left + ((right - left) >> 1)
		if nums[left] <= nums[mid] {
			left = mid
		} else if nums[mid] <= nums[right] {
			right = mid
		}
	}
	return -1
}

/* 中等 */

// 单词接龙（亚马逊在半年内面试常考）
func ladderLength(beginWord string, endWord string, wordList []string) (pathLen int) {
	n := len(beginWord)
	wordListMap := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordListMap[word] = true
	}
	if !wordListMap[endWord] {
		return 0
	}
	visitedFromBegin := map[string]bool{beginWord: true}
	visitedFromEnd := map[string]bool{endWord: true}
	beginQ, endQ := []string{beginWord}, []string{endWord}
	pathLen = 1
	for len(beginQ) > 0 && len(endQ) > 0 {
		if len(beginQ) > len(endQ) { // 选择一个较短的队列进行后续逻辑
			beginQ, endQ = endQ, beginQ
			visitedFromBegin, visitedFromEnd = visitedFromEnd, visitedFromBegin
		}
		pathLen++
		size := len(beginQ)
		var k byte
		for i := 0; i < size; i++ {
			cur := []byte(beginQ[i])
			for j := 0; j < n; j++ {
				old := cur[j]
				for k = 'a'; k <= 'z'; k++ {
					if cur[j] != k {
						cur[j] = k
						nextWod := string(cur)
						if wordListMap[nextWod] && !visitedFromBegin[nextWod] {
							if visitedFromEnd[nextWod] {
								return pathLen
							}
							beginQ = append(beginQ, nextWod)
							visitedFromBegin[nextWod] = true
						}
					}
				}
				cur[j] = old
			}
		}
		beginQ = beginQ[size:]
	}
	return 0

}

// 岛屿数量（近半年内，亚马逊在面试中考查此题达到 350 次）
// 扫雷游戏（亚马逊、Facebook 在半年内面试中考过）
// 跳跃游戏 （亚马逊、华为、Facebook 在半年内面试中考过）
// 搜索旋转排序数组（Facebook、字节跳动、亚马逊在半年内面试常考）
// 搜索二维矩阵（亚马逊、微软、Facebook 在半年内面试中考过）
// 寻找旋转排序数组中的最小值（亚马逊、微软、字节跳动在半年内面试中考过）

/* 困难 */

// 单词接龙 II （微软、亚马逊、Facebook 在半年内面试中考过）
// 跳跃游戏 II （亚马逊、华为、字节跳动在半年内面试中考过）
