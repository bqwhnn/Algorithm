package algorithm

func countingSort(nums []int) {
	len := len(nums)
	min, max := nums[0], nums[0]
	m := make(map[int]int)

	for i := 0; i < len; i++ {
		m[nums[i]]++
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	k := 0
	for i := min; i <= max; i++ {
		for count := m[i]; count > 0; count-- {
			nums[k] = i
			k++
		}
	}
}
