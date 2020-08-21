package main

import "fmt"

func main() {
	// N 皇后（字节跳动、亚马逊、百度在半年内面试中考过）
	fmt.Println(totalNQueens(4)) // 2
}

/* 简单 */
// 位 1 的个数（Facebook、苹果在半年内面试中考过）
func hammingWeight(num uint32) (result int) {
	for num > 0 {
		if ((num - 1) ^ num) == 1 {
			result++
		}
		num >>= 1
	}
	return result
}

// 2 的幂（谷歌、亚马逊、苹果在半年内面试中考过）
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// 颠倒二进制位（苹果在半年内面试中考过）
func reverseBits(num uint32) (result uint32) {
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}

// 数组的相对排序（谷歌在半年内面试中考过）
// 由于数据范围不大，可以采用计数排序
func relativeSortArray(arr1 []int, arr2 []int) (result []int) {
	a1 := make([]int, 1001)
	for _, v := range arr1 {
		a1[v]++
	}
	for _, v := range arr2 {
		for a1[v] > 0 {
			result = append(result, v)
			a1[v]--
		}
	}
	for num, count := range a1 {
		for count > 0 {
			result = append(result, num)
			count--
		}
	}
	return result
}

// 有效的字母异位词（Facebook、亚马逊、谷歌在半年内面试中考过）
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

/* 中等 */
// LRU 缓存机制（亚马逊、字节跳动、Facebook、微软在半年内面试中常考）
// 双向链表+map
type LRUCache struct {
	head, tail *TwoWayListNode         // 从头至尾按从新到旧排序
	index      map[int]*TwoWayListNode // key是关键字
	capacity   int
}

type TwoWayListNode struct {
	Key, Val   int
	Prev, Next *TwoWayListNode
}

func Constructor(capacity int) LRUCache {
	head := &TwoWayListNode{}
	tail := &TwoWayListNode{Prev: head}
	head.Next = tail
	return LRUCache{
		head:     head,
		tail:     tail,
		index:    make(map[int]*TwoWayListNode, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.index[key]; exist {
		// move node to front
		this.moveToFront(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// 如果关键字存在，更新，并移动节点到头部
	if node, exist := this.index[key]; exist {
		node.Val = value
		this.moveToFront(node)
	} else {
		// 关键字不存在，在头部添加节点
		// 添加前确认是否需要还有空余位置，如果没有，删除尾部节点
		if len(this.index) == this.capacity {
			delete(this.index, this.tail.Prev.Key)
			this.removeNode(this.tail.Prev)
		}
		node := &TwoWayListNode{
			Key: key,
			Val: value,
		}
		this.prepend(node)
		this.index[key] = node
	}
}

func (this *LRUCache) removeNode(node *TwoWayListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToFront(node *TwoWayListNode) {
	if node.Prev == this.head {
		return
	}
	this.removeNode(node)
	this.prepend(node)
}

// 从链表头部添加一个节点
func (this *LRUCache) prepend(node *TwoWayListNode) {
	node.Next = this.head.Next
	node.Prev = this.head
	node.Next.Prev = node
	this.head.Next = node
}

// 力扣排行榜（Bloomberg 在半年内面试中考过）
type Leaderboard struct {
	players map[int]int // playerId =>score
	scores  []int       // 排好序的score
}

func Constructor2() Leaderboard {
	return Leaderboard{
		players: make(map[int]int),
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	if _, exist := this.players[playerId]; exist {
		oldScore := this.players[playerId]
		this.players[playerId] += score
		this.deleteScore(oldScore)
		this.insertScore(this.players[playerId])
	} else {
		this.players[playerId] = score
		this.insertScore(score)
	}
}

func (this *Leaderboard) Top(K int) (scores int) {
	i := len(this.scores) - 1
	for K > 0 {
		scores += this.scores[i]
		i--
		K--
	}
	return scores
}

func (this *Leaderboard) Reset(playerId int) {
	this.deleteScore(this.players[playerId])
	delete(this.players, playerId)
}

// score总是存在
func (this *Leaderboard) deleteScore(score int) {
	i := this.search(score)
	this.scores = append(this.scores[:i], this.scores[i+1:]...)
}
func (this Leaderboard) search(score int) (index int) {
	index = -1
	left, right := 0, len(this.scores)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if this.scores[mid] == score {
			return mid
		} else if this.scores[mid] < score {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}
func (this *Leaderboard) insertScore(score int) {
	n := len(this.scores)
	if n == 0 || this.scores[n-1] <= score {
		this.scores = append(this.scores, score)
		return
	} else if this.scores[0] >= score {
		this.scores = append([]int{score}, this.scores...)
		return
	}
	// 找到小于等于score的最大元素
	left, right := 0, n-1
	target := -1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if this.scores[mid] <= score {
			target = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	tmp := make([]int, n+1)
	copy(tmp, this.scores[:target+1])
	tmp[target+1] = score
	copy(tmp[target+2:], this.scores[target+1:])
	this.scores = tmp
}

// 合并区间（Facebook、字节跳动、亚马逊在半年内面试中常考）
func merge(intervals [][]int) (merged [][]int) {
	if len(intervals) <= 1 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	i := 0
	merged = make([][]int, 0, len(intervals))
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		for i < len(intervals)-1 && end >= intervals[i+1][0] {
			i++
			start = getMin(start, intervals[i][0])
			end = getMax(end, intervals[i][1])
		}
		merged = append(merged, []int{start, end})
		i++
	}
	return merged
}
func quickSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(intervals, left, right)
	quickSort(intervals, left, pivot-1)
	quickSort(intervals, pivot+1, right)
}
func partition(intervals [][]int, left, right int) (pivot int) {
	standard := intervals[right]
	pos := left // 小于standard元素要放置的地方
	for i := left; i < right; i++ {
		if intervals[i][0] < standard[0] {
			intervals[pos], intervals[i] = intervals[i], intervals[pos]
			pos++
		}
	}
	intervals[pos], intervals[right] = intervals[right], intervals[pos]
	return pos
}

/* 困难 */
// N 皇后（字节跳动、亚马逊、百度在半年内面试中考过）

// N 皇后 II （亚马逊在半年内面试中考过）
func totalNQueens(n int) (count int) {
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return
	}
	totalNQueensHelper(n, 0, 0, 0, 0, &count)
	return count
}

// 回溯函数
// col, leftDiagonal, rightDiagonal 分别表示在row这一行皇后在垂直、左斜线和右斜线的攻击范围的攻击范围
func totalNQueensHelper(n, row, col, leftDiagonal, rightDiagonal int, count *int) {
	if row == n {
		(*count)++
		return
	}
	// (1 << n) - 1 将n皇后不需要的高位全部赋为0
	available := (^(col | leftDiagonal | rightDiagonal)) & ((1 << n) - 1) // 当前行的可用位置
	for available != 0 {
		pos := available & -available           // 获取最低位的1的位置
		available = available & (available - 1) // 将pos位置置为0，也就是在pos位置放上皇后
		totalNQueensHelper(n, row+1, col|pos, (leftDiagonal|pos)<<1, (rightDiagonal|pos)>>1, count)
	}
}

// 翻转对（字节跳动在半年内面试中考过）
func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return reversePairsSub(nums, 0, len(nums)-1)
}
func reversePairsSub(nums []int, left, right int) (count int) {
	if left >= right {
		return 0
	}
	mid := left + ((right - left) >> 1)
	count = reversePairsSub(nums, left, mid) + reversePairsSub(nums, mid+1, right)
	// 计算两个分区的元素对并合两个并排序区间
	merged := make([]int, right-left+1)
	k := 0
	i := left    // i 代表左区间的下标
	j := mid + 1 // j 代表右区间的下标，用于计算元素对
	q := j       // q 代表右区间的下标，用于合并区间
	for i <= mid {
		// 计算nums[i] 和右侧区间的元素对个数
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		count += j - (mid + 1)
		// 进行合并
		for q <= right && nums[i] >= nums[q] {
			merged[k] = nums[q]
			k++
			q++
		}
		merged[k] = nums[i]
		k++
		i++
	}
	for q <= right {
		merged[k] = nums[q]
		k++
		q++
	}
	copy(nums[left:right+1], merged)
	return count
}

/* 下周预习题目 */
//     不同路径
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

//     最小路径和
func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				if j == 0 {
					dp[j] = grid[i][j]
				} else {
					dp[j] = grid[i][j] + dp[j-1]
				}
			} else {
				if j == 0 {
					dp[j] += grid[i][j]
				} else {
					dp[j] = getMin(dp[j], dp[j-1]) + grid[i][j]
				}
			}
		}
	}
	return dp[len(grid[0])-1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
