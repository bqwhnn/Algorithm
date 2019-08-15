package algorithm

import (
	"fmt"
)

// CountingSort is counting sort
func CountingSort(a []int) []int {
	n := len(a)
	k := a[0]

	for i := 0; i < n; i++ {
		if a[i] > k {
			k = a[i]
		}
	}

	b := make([]int, k+1)
	c := make([]int, n)
	order := make([]int, n)
	for i := 0; i < n; i++ {
		b[a[i]]++
	}

	for i := 1; i <= k; i++ {
		b[i] += b[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		c[b[a[i]]-1] = a[i]
		order[b[a[i]]-1] = i // order 记录了排序之前的位置，表明是稳定排序
		b[a[i]]--
	}

	fmt.Println("counting sort order", order)
	return c
}
