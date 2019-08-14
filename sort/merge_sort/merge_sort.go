package algorithm

func mergeSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)
	sort(nums, temp, 0, n-1)
}

func sort(a, temp []int, first, last int) {
	if first < last {
		mid := (first + last) / 2
		sort(a, temp, first, mid)
		sort(a, temp, mid+1, last)
		merge(a, temp, first, mid, last)
	}
}

func merge(a, temp []int, first, mid, last int) {
	i := first
	j := mid + 1
	k := 0

	for i <= mid && j <= last {
		if a[i] < a[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = a[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = a[i]
		i++
		k++
	}
	for j <= last {
		temp[k] = a[j]
		j++
		k++
	}

	copy(a[first:last+1], temp[0:k])
}
