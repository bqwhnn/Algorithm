package algorithm

import (
	"math"
)

const (
	bc = 10 // 桶的数量
)

// BucketSort is bucket sort
func BucketSort(a []int) []int {
	k, n := a[0], len(a)
	bucket := make([][]int, bc)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		if a[i] > k {
			k = a[i]
		}
	}

	m := int(math.Ceil(float64(k)/bc))
	for i := 0; i < n; i++ {
		bn := a[i] / m // 计算放在哪个桶
		b := bucket[bn]
		b = append(b, a[i])
		for j := len(b)-1; j > 0 && b[j] < b[j-1]; j-- {
			b[j], b[j-1] = b[j-1], b[j]
		}
		bucket[bn] = b
	}

	for i, j := 0, 0; i < bc; i++ {
		b := bucket[i]
		blen := len(b)
		copy(c[j:j+blen], b[0:blen])
		j += blen
	}

	return c
}
