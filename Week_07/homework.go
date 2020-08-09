package main

import "fmt"

func main() {
	// 岛屿数量（近半年内，亚马逊在面试中考查此题达到 361 次）
	fmt.Println(numIslands2([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	// 被围绕的区域（亚马逊、eBay、谷歌在半年内面试中考过）
	// 单词接龙（亚马逊、Facebook、谷歌在半年内面试中考过）
	// 最小基因变化（谷歌、Twitter、腾讯在半年内面试中考过）

	/* 困难 */
	// 解数独（亚马逊、华为、微软在半年内面试中考过）

}

/* 简单 */
// 爬楼梯（阿里巴巴、腾讯、字节跳动在半年内面试常考）
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	pp, p := 1, 1
	for i := 2; i <= n; i++ {
		p, pp = p+pp, p
	}
	return p
}

/* 中等 */
// 实现 Trie (前缀树) （亚马逊、微软、谷歌在半年内面试中考过）

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

// 朋友圈（亚马逊、Facebook、字节跳动在半年内面试中考过）
// 法一：dfs
func findCircleNum(M [][]int) (count int) {
	if len(M) == 0 {
		return 0
	}
	n := len(M)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			findCircleNumDFSHelper(M, visited, i, n)
			count++
		}
	}
	return count
}
func findCircleNumDFSHelper(M [][]int, visited []bool, i, n int) {
	for j := 0; j < n; j++ {
		if M[i][j] == 1 && !visited[j] {
			visited[j] = true
			findCircleNumDFSHelper(M, visited, j, n)
		}
	}
}

// 法二：并查集
func findCircleNum2(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	n := len(M)
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.getCount()
}

type unionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *unionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	return &unionFind{
		parent: p,
		count:  n,
	}
}

// 返回集合的根元素
func (this unionFind) find(p int) int {
	root := p
	for root != this.parent[root] {
		root = this.parent[root]
	}
	// 压缩路径
	for p != this.parent[p] {
		next := this.parent[p]
		this.parent[p] = root
		p = next
	}
	return root
}
func (this *unionFind) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX == rootY {
		return
	}
	this.parent[rootX] = rootY
	this.count--
}
func (this unionFind) getCount() int {
	return this.count
}

// 岛屿数量（近半年内，亚马逊在面试中考查此题达到 361 次）
// 法一：dfs
func numIslands1(grid [][]byte) (count int) {
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

// 法二：并查集
// 「并查集」主要用于解决「动态连通性」问题，重点关注的是连接问题，不关注路径问题。
// 对于本题，就是将水域和周边水域连接，岛屿和周边岛屿连接
// 岛屿的数量就是岛屿联通集合的数目
var around = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func numIslands2(grid [][]byte) (count int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	water := m * n // 归类水域集合
	uf := NewUnionFind(water + 1)
	getIndex := func(x, y int) int {
		return x*n + y
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				uf.union(water, getIndex(i, j))
			} else {
				// 本身是陆地，同时也要合并四周的陆地
				for _, diff := range around {
					newX, newY := i+diff[0], j+diff[1]
					if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == '1' {
						uf.union(getIndex(i, j), getIndex(newX, newY))
					}
				}
			}
		}
	}
	return uf.getCount() - 1
}

// 被围绕的区域（亚马逊、eBay、谷歌在半年内面试中考过）

// 有效的数独（亚马逊、苹果、微软在半年内面试中考过）
// 位运算，使用一个9位的二进制数来判断一个数字是否被使用过，0为未使用，1为已使用
func isValidSudoku(board [][]byte) bool {
	var row, column, squ [9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				squNo := i/3*3 + j/3
				num := int(board[i][j] - '1')
				if isUsedNum(num, row[i]) || isUsedNum(num, squ[squNo]) || isUsedNum(num, column[j]) {
					return false
				}
				row[i] = row[i] ^ (1 << num)
				column[j] = column[j] ^ (1 << num)
				squ[squNo] = squ[squNo] ^ (1 << num)
			}
		}
	}
	return true
}

// set表示9位二进制数，n表示需要判断使用的数字
func isUsedNum(n, set int) bool {
	if ((set >> n) & 1) == 1 {
		return true
	}
	return false
}

// 括号生成（亚马逊、Facebook、字节跳动在半年内面试中考过）
func generateParenthesis(n int) (result []string) {
	if n == 0 {
		return nil
	}
	cur := make([]byte, n*2)
	generateHelper(n, n, cur, 0, &result)
	return result
}
func generateHelper(left, right int, cur []byte, i int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, string(cur))
		return
	}
	if left > 0 {
		cur[i] = '('
		generateHelper(left-1, right, cur, i+1, result)
	}
	if right > left {
		cur[i] = ')'
		generateHelper(left, right-1, cur, i+1, result)
	}
}

// 单词接龙（亚马逊、Facebook、谷歌在半年内面试中考过）
// 最小基因变化（谷歌、Twitter、腾讯在半年内面试中考过）

/* 困难 */
// 单词搜索 II （亚马逊、微软、苹果在半年内面试中考过）
// TRIE树适合查找公共前缀匹配的字符串
// 将 words构建成一棵TRIE树
// 然后使用回溯进行查找
func findWords(board [][]byte, words []string) (result []string) {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	trie := buildTrieTree(words)
	find := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backTracing(board, []byte{}, i, j, m, n, trie, find)
		}
	}
	for k := range find {
		result = append(result, k)
	}
	return
}

var diff [4][2]int = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// 为了去重，结果集使用一个map进行保存
func backTracing(board [][]byte, cur []byte, i, j, m, n int, trie *TrieNode, find map[string]bool) {
	if i < 0 || i == m || j < 0 || j == n || board[i][j] == '.' {
		return
	}
	k := board[i][j] - 'a'
	if trie.next[k] == nil {
		return
	}
	cur = append(cur, board[i][j])
	tmp := board[i][j]
	board[i][j] = '.'
	if trie.next[k].isEnd {
		find[string(cur)] = true
	}
	for _, v := range diff {
		backTracing(board, cur, i+v[0], j+v[1], m, n, trie.next[k], find)
	}
	cur = cur[:len(cur)-1]
	board[i][j] = tmp
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

// N 皇后（亚马逊、苹果、字节跳动在半年内面试中考过）
func solveNQueens(n int) (result [][]string) {
	board := make([][]byte, n)
	tpl := make([]byte, n)
	for i := 0; i < n; i++ {
		tpl[i] = '.'
	}
	for k := range board {
		board[k] = make([]byte, n)
		copy(board[k], tpl)
	}
	nQueensHelper(board, n, 0, &result)
	return result
}
func nQueensHelper(board [][]byte, n, row int, result *[][]string) {
	if row == n {
		tmp := make([]string, n)
		for i, r := range board {
			tmp[i] = string(r)
		}
		*result = append(*result, tmp)
		return
	}
	for col := 0; col < n; col++ {
		if checkPos(board, n, row, col) {
			board[row][col] = 'Q'
			nQueensHelper(board, n, row+1, result)
			board[row][col] = '.'
		}
	}
}
func checkPos(board [][]byte, n, row, colomn int) bool {
	if row == 0 {
		return true
	}
	leftup, rightup := colomn-1, colomn+1
	for i := row - 1; i >= 0; i-- {
		if board[i][colomn] == 'Q' {
			return false
		}
		if leftup >= 0 && board[i][leftup] == 'Q' {
			return false
		}
		if rightup < n && board[i][rightup] == 'Q' {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

// 解数独（亚马逊、华为、微软在半年内面试中考过）
