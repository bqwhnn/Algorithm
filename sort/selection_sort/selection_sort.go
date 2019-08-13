package algorithm

func selectionSort(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
