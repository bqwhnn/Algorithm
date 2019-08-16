package algorithm

// RadixSort is LSD radix sort
func RadixSort(a []int) []int {
	n, m := len(a), a[0]
	divisor := 1

	for i := 0; i < n; i++ {
		if a[i] > m {
			m = a[i]
		}
	}

	for m/divisor > 0 {
		bucket := make([][]int, 10)
		for i := 0; i < n; i++ {
			k := a[i] / divisor % 10
			b := bucket[k]
			b = append(b, a[i])
			bucket[k] = b
		}
		for i, j := 0, 0; i < 10; i++ {
			b := bucket[i]
			for _, v := range b {
				a[j] = v
				j++
			}
		}
		divisor *= 10
	}
	return a
}
