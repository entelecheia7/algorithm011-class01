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
func numIslands1(grid [][]byte) (count int) {
	if len(grid) == 0 {
		return 0
	}
	lx := len(grid)
	ly := len(grid[0])
	visited := make([][]bool, lx)
	for i := 0; i < lx; i++ {
		visited[i] = make([]bool, ly)
	}
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				// 第一次发现一个岛的坐标，递归标记周围是‘1’的坐标为已访问
				visited[i][j] = true
				markIsland(grid, i, j, lx, ly, visited)
				count++
			}

		}
	}
	return count
}

var around = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func markIsland(grid [][]byte, x, y, lx, ly int, visited [][]bool) {
	if x < 0 || y < 0 || x >= lx || y >= ly {
		return
	}
	for _, diff := range around {
		newX, newY := x+diff[0], y+diff[1]
		if newX >= 0 && newX < lx && newY >= 0 && newY < ly && grid[newX][newY] == '1' && !visited[newX][newY] {
			visited[newX][newY] = true
			markIsland(grid, newX, newY, lx, ly, visited)
		}
	}
}

func numIslands2(grid [][]byte) (count int) {
	if len(grid) == 0 {
		return 0
	}
	lx := len(grid)
	ly := len(grid[0])
	visited := make([][]bool, lx)
	for i := 0; i < lx; i++ {
		visited[i] = make([]bool, ly)
	}

	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				queue := [][2]int{{i, j}}
				visited[i][j] = true
				for len(queue) > 0 {
					x, y := queue[0][0], queue[0][1]
					for _, diff := range around {
						newX, newY := x+diff[0], y+diff[1]
						if newX >= 0 && newX < lx && newY >= 0 && newY < ly && grid[newX][newY] == '1' && !visited[newX][newY] {
							visited[newX][newY] = true
							queue = append(queue, [2]int{newX, newY})
						}
					}
					queue = queue[1:]
				}

				count++
			}
		}
	}
	return count
}

// 扫雷游戏（亚马逊、Facebook 在半年内面试中考过）
var relative = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func updateBoard1(board [][]byte, click []int) [][]byte {
	lx := len(board)
	ly := len(board[0])
	queue := [][2]int{{click[0], click[1]}}
	visited := make(map[[2]int]bool, lx*ly) // 防止添加重复的坐标入组，这在大量数据时会造成内存问题
	visited[queue[0]] = true
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		if board[x][y] == 'M' {
			board[x][y] = 'X'
			continue
		}
		nextXY := [][2]int{}
		var mine byte = '0' // 地雷数量
		for _, xy := range relative {
			relativeX, relativeY := x+xy[0], y+xy[1]
			if relativeX < 0 || relativeX >= lx || relativeY < 0 || relativeY >= ly {
				continue
			}
			switch board[relativeX][relativeY] {
			case 'M':
				mine++
			case 'E':
				nextXY = append(nextXY, [2]int{relativeX, relativeY})
			}
		}
		if mine > '0' {
			board[x][y] = mine
		} else {
			board[x][y] = 'B'
			for _, xy := range nextXY {
				if !visited[xy] {
					queue = append(queue, xy)
					visited[xy] = true
				}
			}

		}
	}

	return board
}

// 法二：DFS
func updateBoard2(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	lx := len(board)
	ly := len(board[0])
	updateBoardDFSHelper(board, click[0], click[1], lx, ly)

	return board
}
func updateBoardDFSHelper(board [][]byte, x, y, lx, ly int) {
	if x < 0 || x >= lx || y < 0 || y >= ly || board[x][y] != 'E' {
		return
	}
	var mine byte = '0' // 地雷数量
	board[x][y] = 'B'
	for _, xy := range relative {
		relativeX, relativeY := x+xy[0], y+xy[1]
		if relativeX >= 0 && relativeX < lx && relativeY >= 0 && relativeY < ly && board[relativeX][relativeY] == 'M' {
			mine++
		}
	}
	if mine > '0' {
		board[x][y] = mine
	} else {
		for _, xy := range relative {
			updateBoardDFSHelper(board, x+xy[0], y+xy[1], lx, ly)
		}
	}
}

// 跳跃游戏 （亚马逊、华为、Facebook 在半年内面试中考过）
// 法一：贪心，O(n)
// 从第一个格子逐渐更新可以跳到的最远处
// 直至 到达终点 或 尝试完毕，发现终点不可达
func canJump1(nums []int) (result bool) {
	n := len(nums)
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	furthest := nums[0]
	for i := 0; i <= furthest; i++ {
		furthest = getMax(furthest, i+nums[i])
		if furthest >= n-1 {
			return true
		}
	}
	return false
}

// 法二：动态规划
// 从倒数第二个元素开始判断是否可以到达终点，逐步更新终点
// O(n^2)
func canJump2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	state := make([]bool, n)
	state[n-1] = true
	// 2, 3, 1, 1, 4
	for i := n - 2; i >= 0; i-- {
		// j代表在i处可能走的步数
		for j := 1; j <= nums[i] && i+j < n; j++ {
			if state[i+j] == true {
				state[i] = true
				break
			}
		}
	}
	return state[0]
}

// 搜索旋转排序数组（Facebook、字节跳动、亚马逊在半年内面试常考）
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] { // 左侧为有序数组
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右侧为有序数组
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 搜索二维矩阵（亚马逊、微软、Facebook 在半年内面试中考过）
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[0][0] > target {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	if matrix[m-1][n-1] < target {
		return false
	}
	left, right := 0, m*n-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		midVal := matrix[mid/n][mid%n]
		if midVal == target {
			return true
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 寻找旋转排序数组中的最小值（亚马逊、微软、字节跳动在半年内面试中考过）
// 最小值永远在（中点以）及发生了旋转的一侧
func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	left, right := 0, n-1
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > nums[right] { // 左侧有序，右侧无序
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

/* 困难 */
// 单词接龙 II （微软、亚马逊、Facebook 在半年内面试中考过）
// 双向BFS
func findLadders(beginWord string, endWord string, wordList []string) (result [][]string) {
	if len(wordList) == 0 {
		return nil
	}
	wordListMap := make(map[string]bool, len(wordList)) // 一个全局的未访问元素数组
	for _, w := range wordList {
		wordListMap[w] = true
	}
	if !wordListMap[endWord] {
		return nil
	}
	delete(wordListMap, endWord)

	// queue保存这一层的节点
	queue, queueFromEnd := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	n := len(beginWord)
	var j byte
	endFlag, reverseFlag := false, false
	path := make(map[string][]string) // 记录key可以转换的单次

	for len(queue) > 0 && len(queueFromEnd) > 0 && !endFlag {
		if len(queue) > len(queueFromEnd) {
			queue, queueFromEnd = queueFromEnd, queue
			reverseFlag = !reverseFlag
		}
		for w := range queue {
			delete(wordListMap, w)
		}
		newqueue := make(map[string]bool)
		for word := range queue {
			tmp := []byte(word)
			for i := 0; i < n; i++ {
				old := tmp[i]
				for j = 'a'; j <= 'z'; j++ {
					if j != old {
						tmp[i] = j
						convertion := string(tmp)
						if queueFromEnd[convertion] { // 双向BFS相遇
							if reverseFlag {
								path[convertion] = append(path[convertion], word)
							} else {
								path[word] = append(path[word], convertion)
							}
							endFlag = true
						} else if wordListMap[convertion] { // 未访问过，说明到达下一层
							newqueue[convertion] = true
							if reverseFlag {
								path[convertion] = append(path[convertion], word)
							} else {
								path[word] = append(path[word], convertion)
							}
						}
					}
				}
				tmp[i] = old
			}
		}
		queue = newqueue
	}

	// DFS，从beginWord开始组装结果
	cur := []string{beginWord}
	var generator func([]string)
	generator = func(words []string) {
		for _, n := range words {
			cur = append(cur, n)
			if n == endWord {
				tmp := make([]string, len(cur))
				copy(tmp, cur)
				result = append(result, tmp)
			} else {
				generator(path[n])
			}
			cur = cur[:len(cur)-1]
		}

	}
	generator(path[beginWord])

	return result
}

// 跳跃游戏 II （亚马逊、华为、字节跳动在半年内面试中考过）
func jump(nums []int) (step int) {
	n := len(nums)
	if n == 1 {
		return 0
	}
	// 因为已经确定终点可达，所以可以直接遍历
	// 只要达到 n-2处，必定可达到终点
	maxPos := 0
	reachable := 0
	for i := 0; i < n-1; i++ {
		maxPos = getMax(maxPos, i+nums[i])
		if i == reachable {
			reachable = maxPos
			step++
		}
	}
	return step
}
