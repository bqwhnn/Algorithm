package algorithm

// HeapSort is heap sort
func HeapSort(a []int) {
	n := len(a)
	makeHeap(a, n)
	
	for i := n-1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		// n 个元素的堆，将堆首与堆尾对调，堆元素变成 n-1 个
		fixDown(a, 0, i)
	}
}

// 建立最大堆
func makeHeap(a []int, n int) {
	for i := n/2; i >= 0; i--{
		fixDown(a, i, n)
	}
}

// 添加元素, k 是要添加的元素
func pushHeap(a []int, n, k int) {
	a[n] = k
	fixUp(a, n)
}

// 删除堆尾元素
func popHeap(a []int, n int) {
	a[0], a[n] = a[n], a[0]
	fixDown(a, 0, n-1)
}

// 向下调整堆结点, i 是向下调整的结点序号，n 是数组元素个数
func fixDown(a []int, i, n int) {
	j := 2 * i + 1
	for j < n {
		if j+1 < n && a[j+1] > a[j] {
			j++
		}

		if a[j] < a[i] {
			break
		}

		a[i], a[j] = a[j], a[i]
		i = j
		j = 2 * i + 1
	}
}

// 向上调整堆结点, i 是向上调整的结点序号
func fixUp(a []int, i int) {
	j := i / 2
	for j >= 0 && i != 0 && a[j] < a[i] {
		a[i], a[j] = a[j], a[i]
		i = j
		j = i / 2
	}
}
