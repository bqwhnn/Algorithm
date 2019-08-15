# algorithm
Data Structure and Algorithm in Go

* 数据结构
  * 数组
  * 栈
  * 队列
  * 链表
  * 哈希表
  * 二叉树
  * 堆
* 算法
  * 排序
    * [冒泡排序](https://github.com/bqwhnn/algorithm/tree/master/sort/bubble_sort)
    * [选择排序](https://github.com/bqwhnn/algorithm/tree/master/sort/selection_sort)
    * [插入排序](https://github.com/bqwhnn/algorithm/tree/master/sort/insertion_sort)
    * [希尔排序](https://github.com/bqwhnn/algorithm/tree/master/sort/shell_sort)
    * [归并排序](https://github.com/bqwhnn/algorithm/tree/master/sort/merge_sort)
    * [快速排序](https://github.com/bqwhnn/algorithm/tree/master/sort/quick_sort)
    * [堆排序](https://github.com/bqwhnn/algorithm/tree/master/sort/heap_sort)
    * [计数排序](https://github.com/bqwhnn/algorithm/tree/master/sort/counting_sort)
    * [桶排序](https://github.com/bqwhnn/algorithm/tree/master/sort/bucket_sort)
    * [基数排序](https://github.com/bqwhnn/algorithm/tree/master/sort/radix_sort)


排序算法 | 平均时间复杂度 | 最佳情况 | 最坏情况 | 空间复杂度 | 排序方式 | 稳定性
:-: | :-: | :-: | :-: | :-: | :-: | :-: 
冒泡排序 | 
选择排序 |
插入排序 |
希尔排序 |
归并排序 |
快速排序 |
堆排序 |
计数排序 |
桶排序 |
基数排序 |

冒泡排序、选择排序、插入排序、希尔排序、归并排序、快速排序、堆排序需要对系列中的数据进行比较，所以称为基于比较的排序。

基于比较的排序算法的时间复杂度是不能突破 O(nlogn) 的。

> N个数有N!个可能的排列情况，也就是说基于比较的排序算法的判定树有N!个叶子结点，比较次数至少为log(N!)=O(NlogN)(斯特林公式)。

计数排序、桶排序、基数排序是非基于比较的排序。

> 非基于比较的排序算法的使用都是有条件限制的，例如元素的大小限制。元素的个数决定了桶的规模，影响算法的效率。所以对于特定场合有着特殊的性质数据，非基于比较的排序算法则能够非常巧妙地解决。
