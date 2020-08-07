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

	// 岛屿数量
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
			findWordsBackTracing(board, used, []byte{}, i, j, m, n, trie, find)
		}
	}
	for k := range find {
		result = append(result, k)
	}

	return
}

// 为了去重，结果使用一个map
func findWordsBackTracing(board [][]byte, used [][]bool, cur []byte, i, j, m, n int, trie *TrieNode, find map[string]bool) {
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
		findWordsBackTracing(board, used, cur, i-1, j, m, n, trie.next[k], find)
	}
	if i < m-1 && !used[i+1][j] {
		findWordsBackTracing(board, used, cur, i+1, j, m, n, trie.next[k], find)
	}
	if j > 0 && !used[i][j-1] {
		findWordsBackTracing(board, used, cur, i, j-1, m, n, trie.next[k], find)
	}
	if j < n-1 && !used[i][j+1] {
		findWordsBackTracing(board, used, cur, i, j+1, m, n, trie.next[k], find)
	}
	cur = cur[:len(cur)-1]
	used[i][j] = false
}

func buildTrieTree(words []string) *TrieNode {
	trie := &TrieNode{
		char: '/',
	}
	for _, word := range words {
		p := trie
		lastIndex := len(word) - 1
		for i := 0; i <= lastIndex; i++ {
			k := word[i] - 'a'
			if p.next[k] == nil {
				p.next[k] = &TrieNode{
					char: word[i],
				}
			}
			if i == lastIndex {
				p.next[k].isEnd = true
			} else {
				p = p.next[k]
			}

		}
	}
	return trie
}

type TrieNode struct {
	char  byte
	isEnd bool
	next  [26]*TrieNode
}

// 岛屿数量
// 有效的数独
// N 皇后
// 单词接龙
// 二进制矩阵中的最短路径
