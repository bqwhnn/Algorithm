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


排序算法 | 平均时间复杂度 | 最佳情况 | 最坏情况 | 空间复杂度 | 稳定性
:-: | :-: | :-: | :-: | :-: | :-: 
冒泡排序 | O(n²) | O(n) | O(n²) | O(1) | 稳定
选择排序 | O(n²) | O(n²) | O(n²) | O(1) | 不稳定
插入排序 | O(n²) | O(n) | O(n²) | O(1) | 稳定
希尔排序 | O(nlog^2n) | O(nlogn) | O(n²) | O(1) | 不稳定
归并排序 | O(nlogn) | O(nlogn) | O(nlogn) | O(n) | 稳定
快速排序 | O(nlogn) | O(nlogn) | O(n²) | O(logn) | 不稳定
堆排序 | O(nlogn) | O(nlogn) | O(nlogn) | O(1) | 不稳定
计数排序 | O(n+k) | O(n+k) | O(n+k) | O(n+k) | 稳定
桶排序 | O(n+nlog(n/k)) | O(n) | O(n+nlog(n/k)) | o(n+k) | 稳定
基数排序 | | | | O(n+k) | 稳定

冒泡排序、选择排序、插入排序、希尔排序、归并排序、快速排序、堆排序需要对系列中的数据进行比较，所以称为基于比较的排序。

基于比较的排序算法的时间复杂度是不能突破 O(nlogn) 的。

> N个数有N!个可能的排列情况，也就是说基于比较的排序算法的判定树有N!个叶子结点，比较次数至少为log(N!)=O(NlogN)(斯特林公式)。

计数排序、桶排序、基数排序是非基于比较的排序。

> 非基于比较的排序算法的使用都是有条件限制的，例如元素的大小限制。元素的个数决定了桶的规模，影响算法的效率。所以对于特定场合有着特殊的性质数据，非基于比较的排序算法则能够非常巧妙地解决。

> 用逆序数来理解，假设我们要从小到大排序，一个数组中取两个元素如果前面比后面大，则为一个逆序，容易看出排序的本质就是消除逆序数，可以证明对于随机数组，逆序数是 O(N^2) 的，而如果采用“交换相邻元素”的办法来消除逆序，每次正好只消除一个，因此必须执行 O(N^2) 的交换次数，这就是为啥冒泡、插入等算法只能到平方级别的原因，反过来，基于交换元素的排序要想突破这个下界，必须执行一些比较，交换相隔比较远的元素，使得一次交换能消除一个以上的逆序，希尔、快排、堆排等等算法都是交换比较远的元素，只不过规则各不同罢了
