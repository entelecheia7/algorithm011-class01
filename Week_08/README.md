# 位运算
- 左移（<<）
- 右移（>>）
- 或（|）：二进制位均为0则为0，其他情况为1
- 与（&）：二进制位均为1则为1，其他情况是0
- 取反（~）：二进制位反转，1变0，0变1
- 异或（^）：二进制位相同是0，不同是1

异或操作的特点：
```
x ^ 0 = x
x ^ (~0) = ~x
x ^ (~x) = ~0
x ^ x = 0
c = a ^ b  =>  a ^ c = b, b ^ c = a
a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c
```
常用的位运算：
- 将x右n位清零：x & (~0 << n)
- 获取x第n位的值：(x >> n) & 1
- 获取x第n位的幂值：x & (1<<n)
- 将第n位置为1：x | (1<<n)
- 将第n位置为0：x & (~(1<<n))
- 将x最高位至第n位（含）清零：x & ((1<<n)-1)
- 判断奇偶：
    - (x & 1) == 1
    - (x & 1) == 0
- 清零最低位的1：x & (x-1)
- 得到最低位的1：x & -x
- 判等：x & (~x) == 0


# 布隆过滤器
布隆过滤器由一个很长的二进制向量和一系列随机映射函数组成。它用于检测一个元素是否在一个集合中。它的优点在于空间消耗少、查询时间少，但有一定的误识别率、删除困难。  
布隆过滤器的应用场景有：数据库缓存层、垃圾邮件过滤、推荐去重、分布式系统判断数据是否存在于当前节点。

# 排序算法
排序大致分为两类，比较类排序和非比较类排序。  
比较类排序通过比较决定元素间的相对次序，时间复杂度最好是 O(logn)，也称为非线性时间比较类排序。  
非比较类排序不通过比较来决定元素间的相对次序，时间复杂度是线性的，因此也称为线性时间比较类排序。  
比较类排序包括：
- 交换排序
    - 冒泡排序
    - 快速排序
- 插入排序
    - 插入排序
    - 希尔排序
- 选择排序
    - 选择排序
    - 堆排序
- 归并排序
    - 二路归并排序
    - 多路归并排序

非比较类排序包括：
- 计数排序
- 桶排序
- 基数排序

## 选择排序
每次找最小值，放到待排序数组的起始位置。
```
func selectionSort(nums []int, n int) {
	pos := 0 // 待排序数组起始位置
	for pos < n {
		minest := pos
		for i := pos + 1; i < n; i++ {
			if nums[i] < nums[minest] {
				minest = i
			}
		}
		nums[pos], nums[minest] = nums[minest], nums[pos]
		pos++
	}
}
```

## 插入排序
对于未排序数据，在已排序序列中从后往前扫描，找到相应位置插入。
```
func insertionSort(nums []int, n int) {
	pos := 1 // 未排序数据的起始位置
	for pos < n {
		i := pos - 1
		cur := nums[pos] // 本次循环的待插入数据
		for i >= 0 && nums[i] > cur {
			nums[i+1] = nums[i]
			i--
		}
		nums[i+1] = cur
		pos++
	}
}
```

## 冒泡排序
多层循环，每次检查相邻元素的顺序，逆序则交换。
```
func BubbleSort(nums []int, n int) {
    for j := 0; j < n-1; j++ {
        flag := true
        for i := 0; i < n-1; i++ {
            if nums[i] > nums[i+1] {
                falg = false
                nums[i], nums[i+1] = nums[i+1], nums[i]
            }
        }
        if flag {
            break
        }
    }
}

```

## 快速排序
取一个基准元素，将小元素放在基准左侧，大元素放右侧，然后递归地对两侧数组进行快排。
```
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}
func partition(nums []int, left, right int) (pivot int) {
	standard := nums[right]
	smaller := left // 小于standard的元素的放置位置
	for i := left; i < right; i++ {
		if nums[i] < standard {
			nums[i], nums[smaller] = nums[smaller], nums[i]
			smaller++
		}
	}
	nums[smaller], nums[right] = nums[right], nums[smaller]
	return smaller
}
```

## 归并排序
把数组一分为二，对两个子数组采用归并排序，再将排序好的子数组进行合并。
```
func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + ((right - left) >> 1)
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}
func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j := left, mid+1
	k := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	for a := 0; a < len(tmp); a++ {
		nums[left+a] = tmp[a]
	}
    // 或使用 copy(nums[left:right+1], tmp)
}
```

## 堆排序
构建一个小顶堆或大顶堆，再依次取堆顶元素
```
func heapSort(nums []int, n int) {
	if n == 0 {
		return
	}
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	for j := n - 1; j >= 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		heapify(nums, j, 0)
	}
}

// 大顶堆，自下而上堆化
func heapify(nums []int, n, i int) {
	for {
		maxPos := i
		// 左子节点
		left := 2*i + 1
		if left < n && nums[maxPos] < nums[left] {
			maxPos = left
		}
		// 右子节点
		right := left + 1
		if right < n && nums[maxPos] < nums[right] {
			maxPos = right
		}
		if maxPos == i {
			break
		}
		nums[i], nums[maxPos] = nums[maxPos], nums[i]
		i = maxPos
	}
}
```

## 计数排序
将待排序数组使用额外的数组统计每个元素出现的频次，再按照顺序输出。
这要求待排序数据是整数，且范围不大。

## 桶排序
计数排序的升级版。将数据划分到数量有限的桶，对每个桶的数据进行排序，最后进行合并。

## 基数排序
取得数组的最大位数。按照位来排序，先排个位，再排十位，以此类推。
这要求待排序数据是整数。