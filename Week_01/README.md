# 作业
1. 用 add first 或 add last 这套新的 API 改写 Deque 的代码
```
Deque<String> deque = new LinkedList<String>();
deque.addFirst("a");
deque.addFirst("b");
deque.addFirst("c");
System.out.println(deque);
String str = deque.peekFirst();
System.out.println(str);
System.out.println(deque);
while (deque.size() > 0) {
 System.out.println(deque.removeFirst());
}
System.out.println(deque);

```

2. 分析 Queue 的源码
Queue 是一个接口。它提供了两组方法，每一组方法都包括三个方法：添加元素、弹出队头元素和查看队头元素。不同组的方法的错误处理方式不同，add、remove、element 方法在操作失败时抛出异常，offer、poll、peek 方法在操作失败时返回特殊值。
Queue 的实现有很多种，包括 AbstractQueue<E>, ArrayBlockingQueue<E>, ArrayDeque<E>, ConcurrentLinkedQueue<E>, DelayQueue<E,extends,Delayed>, LinkedBlockingDeque<E>, LinkedBlockingQueue<E>, LinkedList<T>, PriorityBlockingQueue<E>, PriorityQueue<E>, SynchronousQueue<E>。
&nbsp; 
3. 分析 Priority Queue 的源码
Priority Queue 是一个类。它实现了 Queue 接口，此外还提供了方法：clear(清空优先队列)、comparator(查看优先级比较类)、contains、toArray 等常用方法。
优先队列不支持 null 元素，没有容量限制。它基于一个堆的数组实现，默认按照自然顺序排列，或者在构建传入一个优先级比较器。
offer() 的源码实现：
    1. 检查元素是否为空
    2. 增加modCount
    3. 元素个数达到容量则进行扩容
    4. 队列为空时直接将元素放到 index 为 0 的位置
    5. 将元素放到数组的末尾元素之后，再进行堆化（将新元素上浮到合适的位置）。
    
    自下而上堆化堆化时，使用 `parent = (k - 1) >>> 1` 计算父元素位置，不断与父元素比较大小，如果比父元素小，就交换位置。
    poll() 的源码实现：
	1. 检查队列是否为空
    2. 减少元素数量，增加modCount
    3. 取出队首元素
    4. 将队末元素（优先级最大）放到队首，再进行堆化（将队首元素下沉到合适的位置）。  
    
    自上而下堆化时，使用 `int child = (k << 1) + 1` 和 `right = child + 1` 计算左右子元素位置，如果比子元素大，就和两个子节点的较小者交换位置。
    modCount 是用于实现 fast-fail 特性的。 fast-fail 指的是在操作中的多个点检查系统状态，立即报告故障。在这里是用于迭代在迭代时检查 modCount 是否被修改，如果被修改就抛出异常。


# 要点
数组：连续内存，插入、删除O(n)，随机访问O(1)，查询O(n)，无法动态扩容，利于CPU预读
链表：零散内存，插入、删除O(1)，随机访问O(n)，查询O(n)，可动态扩容，不利于CPU预读
跳表：基于有序数据，时间复杂度O(logn)，空间复杂度O(n)
栈：后进先出，先进后出，插入、删除O(1)
队列：先进先出，后进后出，插入、删除O(1)
- 优先队列：可以用堆实现，和堆的时间复杂度相同，插入、删除O(logn)
- 双端队列：插入、删除O(1)