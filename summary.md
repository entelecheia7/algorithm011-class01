## 数据结构
### 基础脉络
- 一维：
    - 基础：数组 array (string), 链表 linked list
    - 高级：栈 stack, 队列 queue, 双端队列 deque, 集合 set, 映射 map (hash or map), etc
- 二维：
    - 基础：树 tree, 图 graph
    - 高级：二叉搜索树 binary search tree (red-black tree, AVL), 堆 heap, 并查集 disjoint set, 字典树 Trie, etc
- 特殊：
    - 位运算 Bitwise, 布隆过滤器 BloomFilter
    - LRU Cache

### Tips
1. 数据结构优化的思路主要是两个：
    - 空间换时间
    - 一维的数据结构，如果要优化，经常采用的方式是升维，也就是说，升成二维。多一个维度，就多一级信息。比如链表优化为跳表。

2. 链表是一种特殊的树，树是一种特殊的图。没有环的图就是树。


## 算法
### 基础脉络
- 基础：
    - If-else, switch —> branch
    - for, while loop —> Iteration
    - 递归 Recursion (Divide & Conquer, Backtrace)
- 搜索 Search: 深度优先搜索 Depth fifirst search, 广度优先搜索 Breadth fifirst search, A*, etc
- 动态规划 Dynamic Programming
- 二分查找 Binary Search
- 贪心 Greedy
- 数学 Math , 几何 Geometry

所有的高级算法都是基于基础的 IF/ELSE、循环和递归的。  
高级算法的根本就是找到它的**重复单元**，转化为基础语句。

### 常见的时间复杂度及对应算法
- O(logn)：二分查找，斐波那切数列求第n项
- O(n)：二叉树的前序&中序&后序遍历、图的遍历、DFS深度优先、BDF广度优先
- O(n^2)：冒泡、插入、选择排序
- O(nlogn)：快排，归并排序、堆排序