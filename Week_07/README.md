# Trie树
Trie树常用于统计和排序大量字符串，适用于自动输入补全、屏蔽字检测等场景。  
在叶子结点可以存储单词的频次，以优化自动输入补全等功能。  
字符集比较小、字符串前缀重合多的情况比较适合使用trie树结构。  

单词搜索2的解题步骤和时间复杂度如下：
1. 根据输入的字符串列表words构建trie树。需要遍历所有字符串，以n表示words数量，k表示字符串平均长度，这部分的时间复杂度是O(kn)。
2. 根据board进行回溯查找，这部分的时间复杂度取决于树的深度，以k表示words的平均长度，以m表示board中的字符个数，则这部分的时间复杂度是O(km)。
综上，整体的时间复杂度为O(k*(m+n))。

# 并查集
并查集是一种主要用于解决“动态连通性”问题的数据结构，用来解决快速判断两个元素是否在同一个集合的问题。比如leetcode的岛屿数量问题。  
并查集支持以下基本操作：
- 创建一个并查集，其中包含n个单元素集合。
- 合并元素x和元素y的集合，前提是集合x和集合y不相交，相交则不需要合并。
- 找到元素x所在的集合。该操作可以用于判断两个元素是否位于一个合集。
它在数据结构上是一颗树。 

代码模板如下：
```
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
```

# 搜索
搜索算法优化有两个方向：
- 避免重复的中间结果（斐波那契数列计算问题）
    - 保存重复的中间结果 或 直接使用递推避免重复的中间结果
- 剪枝：去掉重复分枝或次优分枝

除了这两个方向，还有更高级的搜索算法：
- 双向搜索
- 启发式搜索（A*算法）：根据优先队列进行搜索
## 双向BFS
代码模板：
```
// 1.定义beginQ、endQ
// 2.定义beginVisited、endVisited
// 3.开始搜索
for len(beginQ) > 0 && len(endQ) > 0 {
	// 选择更少的数量进行搜索
	if len(beginQ) > len(endQ) {
		beginQ, endQ = endQ, beginQ
		beginVisited, endVisited = endVisited, beginVisited
	}
	size := len(beginQ)
	for i:= 0; i < size;i++ {
		// 处理当前条目，得到下一层
		……
		if !beginVisited[next] {
			if endVisited[next] {
				// 处理结果，返回
				return 
			}
			beginQ = append(beginQ, next)
			beginVisited[next] = true
		}
	}

	// reverse
	// 根据具体题目这一步骤可以省略
	beginQ, endQ = endQ, beginQ
	beginVisited, endVisited = endVisited, beginVisited
}

```
## A*搜索
A* 搜索的关键在于估价函数 h(n)，它用于评价哪些结点最可能是我们想找到的结点。h(n) 返回一个非负实数，表示从结点 n 到模板结点路径的估计成本，值越大说明优先级越大。  
代码模板：
```
func astarSearch(graph [][]int, start, end int) {
	pq := newPriorityQueue()
	pq.add([2]int{0, 0})
	// init visited
	……
	visited[0][0] = true

	for pq.Len() > 0 {
		cur := pq.pop()
		visited[cur] = true
		nextStep := getNextStep(graph, cur, visited)
		for _, next := range nextStep {
			pq.push(next)
		}
	}
}
```
# AVL树
AVL树种引入了一个平衡因子（Balance Factor）的概念，表示左子树和右子树的高度差，平衡因子的取值范围是 -1、0、1。它始终保证任何一个节点的平衡因子都在取值范围内。  
AVL 树需要存储格外的平衡因子信息，且调整次数频繁，开销比较大。

# 红黑树
红黑树是一种近似平衡的二叉搜索树，它可以保证任意结点的左右子树高度差小于两倍。  
红黑树中的节点，一类被标记为黑色，一类被标记为红色。同时满足以下要求：
- 根节点是黑色的
- 每个叶子节点都是黑色的空节点（NIL），也就是说，叶子节点不存储数据
- 任何相邻的节点都不能同时为红色，也就是说，红色节点是被黑色节点隔开的
- 每个节点，从该节点到达其可达叶子节点的所有路径，都包含相同数目的黑色节点