package algorithm

// InsertionSort is insertion sort
func InsertionSort(a []int) {
	len := len(a)
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0 && a[j] > a[j+1]; j-- {
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
}

// InsertionSort2 减少了交换操作
func InsertionSort2(a []int) {
	len := len(a)
	for i := 1; i < len; i++ {
		j := i - 1
		v := a[i]
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
}
