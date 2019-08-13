package algorithm

func insertionSort(nums []int) {
	len := len(nums)
	for i := 1; i < len; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
