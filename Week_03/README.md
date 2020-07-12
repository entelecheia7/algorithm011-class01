学习笔记
# 代码模板
```
// 递归模板
func recursion(level, maxLevel, params..., result) {
    // recursion terminator
    if level > maxLevel {
        processResult
        return
    }
    // process logic in current level
    process current level

    // drill down
    recursion(level+1, maxLevel, params..., result)

    // reverse current status if needed
    reverse
} 
// 分治代码模板
func divideConquer(problem, param) {
    // terminator
    if problem == nil {
        return result
    }
    // divide problem
    subproblems := splitProbliem(problem, param)
    // conquer subproblems
    subresult1 := divideConquer(subproblems[0], param1)
    subresult2 := divideConquer(subproblems[1], param2)
    ……
    // merge result
    result = mergeResult(subresult1, subresult2...)
    // revert status if needed
}
```

# 代码优化思路
1. 避免数组的扩容操作
2. 找出题目的隐含条件，提前剪枝

# 要点
1. 二叉搜索树的中序遍历是递增，利用这一点可以解决一部分二叉搜索树的题目。
2. 二叉搜索树的节点大小不仅要看根节点的左右子节点，而是要看整个左右子树。