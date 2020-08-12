package main

import (
	"fmt"
)

func main() {
	// 单词搜索 II
	fmt.Println(findWords([][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}, []string{"aaaaaaaaaaaa", "aaaaaaaaaaaaa", "aaaaaaaaaaab"})) // ["aaaaaaaaaaaa"]

	// 有效的数独
	// N 皇后
	// 单词接龙
}

// 实现 Trie (前缀树)
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, char := range word {
		c := char - 'a'
		if this.next[c] == nil {
			this.next[c] = &Trie{}
		}
		this = this.next[c]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, char := range word {
		c := char - 'a'
		if this.next[c] == nil {
			return false
		}
		this = this.next[c]
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, char := range prefix {
		c := char - 'a'
		if this.next[c] == nil {
			return false
		}
		this = this.next[c]
	}
	return true
}

// 单词搜索 II
// TRIE树适合查找公共前缀匹配的字符串
// 将 words构建成一棵TRIE树
// 然后使用回溯进行查找
func findWords(board [][]byte, words []string) (result []string) {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	trie := buildTrieTree(words)
	used := make([][]bool, m)
	for k := range used {
		used[k] = make([]bool, n)
	}
	find := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backTracing(board, used, []byte{}, i, j, m, n, trie, find)
		}
	}
	for k := range find {
		result = append(result, k)
	}

	return
}

// 为了去重，结果使用一个map
func backTracing(board [][]byte, used [][]bool, cur []byte, i, j, m, n int, trie *TrieNode, find map[string]bool) {
	k := board[i][j] - 'a'
	if trie.next[k] == nil {
		return
	}
	cur = append(cur, board[i][j])
	used[i][j] = true
	if trie.next[k].isEnd {
		find[string(cur)] = true
	}

	if i > 0 && !used[i-1][j] {
		backTracing(board, used, cur, i-1, j, m, n, trie.next[k], find)
	}
	if i < m-1 && !used[i+1][j] {
		backTracing(board, used, cur, i+1, j, m, n, trie.next[k], find)
	}
	if j > 0 && !used[i][j-1] {
		backTracing(board, used, cur, i, j-1, m, n, trie.next[k], find)
	}
	if j < n-1 && !used[i][j+1] {
		backTracing(board, used, cur, i, j+1, m, n, trie.next[k], find)
	}
	cur = cur[:len(cur)-1]
	used[i][j] = false
}

func buildTrieTree(words []string) *TrieNode {
	trie := &TrieNode{}
	for _, word := range words {
		p := trie
		for _, char := range word {
			c := char - 'a'
			if p.next[c] == nil {
				p.next[c] = &TrieNode{}
			}
			p = p.next[c]
		}
		p.isEnd = true
	}
	return trie
}

type TrieNode struct {
	isEnd bool
	next  [26]*TrieNode
}

// 岛屿数量
// 法一：dfs
func numIslands(grid [][]byte) (count int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(grid))
	for k := range visited {
		visited[k] = make([]bool, len(grid[0]))
	}
	for i, row := range grid {
		for j, v := range row {
			if !visited[i][j] && v == '1' {
				// 第一次发现一个岛的坐标，递归标记周围是‘1’的坐标为已访问
				markIsland(grid, i, j, visited)
				count++
			}
		}
	}
	return count
}

func markIsland(grid [][]byte, x, y int, visited [][]bool) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || visited[x][y] || grid[x][y] == '0' {
		return
	}
	visited[x][y] = true
	markIsland(grid, x+1, y, visited)
	markIsland(grid, x-1, y, visited)
	markIsland(grid, x, y+1, visited)
	markIsland(grid, x, y-1, visited)
}

// 有效的数独
// N 皇后
// 单词接龙

// 二进制矩阵中的最短路径
// 法一：BFS
var around = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

// 省略了visied数组，直接在grid上进行修改
func shortestPathBinaryMatrix(grid [][]int) (level int) {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	} else if n <= 2 {
		return n
	}
	queue := [][2]int{{0, 0}}
	grid[0][0] = 2
	level++
	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, diff := range around {
				x, y := cur[0]+diff[0], cur[1]+diff[1]
				if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 0 {
					if x == n-1 && y == n-1 {
						return
					}
					queue = append(queue, [2]int{x, y})
					grid[x][y] = 2
				}
			}

		}
		queue = queue[size:]
	}
	return -1
}

// 法二：A*
// 估价函数h(n)代表从当前点到终点的曼哈顿距离（坐标差绝对值之和）
// 优先级是估价函数的值加上当前点已走的距离，只使用曼哈顿距离只能得到一个较优值
// 使用一个小顶堆取代pq
func shortestPathBinaryMatrix2(grid [][]int) (minDist int) {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	} else if n <= 2 {
		return n
	}
	var pq priorityQueue
	maxPos := n - 1
	pq = append(pq, node{x: 0, y: 0})
	dist := make(map[[2]int]int, n*n)
	dist[[2]int{0, 0}] = 1
	for len(pq) > 0 {
		cur := heap.Pop(&pq).(node)
		if grid[cur.x][cur.y] == 2 {
			continue
		}
		if cur.x == maxPos && cur.y == maxPos {
			return dist[[2]int{maxPos, maxPos}]
		}
		grid[cur.x][cur.y] = 2
		for _, diff := range around {
			x, y := cur.x+diff[0], cur.y+diff[1]
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] != 1 {
				heap.Push(&pq, node{x, y, heuristicHelper(x, y, maxPos) + dist[[2]int{cur.x, cur.y}] + 1})
				// 剪枝，同一个点有多条到达路径，如果有更短的路径，就更新
				if dist[[2]int{x, y}] == 0 || dist[[2]int{cur.x, cur.y}]+1 < dist[[2]int{x, y}] {
					dist[[2]int{x, y}] = dist[[2]int{cur.x, cur.y}] + 1
				}
			}
		}
	}
	return -1
}

type priorityQueue []node
type node struct {
	x, y     int
	priority int
}

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(node)) }
func (pq *priorityQueue) Pop() interface{} {
	last := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return last
}

// 估价函数取曼哈顿距离
// 这里做了一个优化，直接取最大的边长
// 返回值越低越好
func heuristicHelper(i, j, maxPos int) int {
	return getMax((maxPos - i), (maxPos - j))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
