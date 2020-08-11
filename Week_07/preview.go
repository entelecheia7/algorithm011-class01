package main

import "fmt"

func main() {
	// fmt.Println(isAnagram("rat", "car")) // false

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // -1
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}

// LRU 缓存机制
//  运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
// 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
// 进阶:
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？
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

// 有效的字母异位词
// 使用哈希表或数组判断相同元素的个数
// 如果有unicode字符使用哈希表，只有英文小写字母可使用数组
func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	var freq [26]int
	for i := 0; i < ls; i++ {
		if s[i] != t[i] {
			freq[s[i]-'a']++
			freq[t[i]-'a']--
		}
	}
	for i := 0; i < 26; i++ {
		if freq[i] != 0 {
			return false
		}
	}
	return true
}
