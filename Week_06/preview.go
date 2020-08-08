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
	// 二进制矩阵中的最短路径

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
