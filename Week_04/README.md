# 卡特兰公式
令h(0)=1,h(1)=1，  
卡塔兰数的递推式为：h(n) = h(0)*h(n-1) + h(1)*h(n-2) + ... + h(n-1)h(0)（n>=2）  
化简为：h(n) = h(n-1)*(4*n-2) / (n+1)

# 代码模板
```
// DFS深度优先
// 递归版
func dfs(node *TreeNode, visited []*TreeNode) {
    if node == nil {
        return
    }
    if inVisited(node) {
        return
    }
    visited = append(visited, node)
    // process current node
    doSomething()

    for _, child := range node.Children {
        if !inVisited(child) {
            dfs(child, visited)
        }
    }
}
// 迭代版
func dfs(node *TreeNode, visited []*TreeNode) {
    if node == nil {
        return
    }
    stack := []*TreeNode{node}
    visited := []*TreeNode{}
    for len(stack) > 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // process current node
        doSomething()
        if inVisited(cur) {
            continue
        }
        visited = append(visited, cur)
        for _, child := range cur.Children {
            stack = append(stack, child)
        }
    }
}

// BFS
// 在图的算法中要加上 visited
func levelOrder(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
    visited := []*Node{}
	for len(queue) != 0 {
		n := len(queue) // 本层节点数量
		val := []int{}  // 本层节点值
		for i := 0; i < n; i++ {
			val = append(val, queue[i].Val)
            // visited = append(visited, queue[i])
			if queue[i].Children != nil {
				queue = append(queue, queue[i].Children...)
			}
		}
		queue = queue[n:]
		result = append(result, val)
	}

	return
}
```