package algorithm

func insertionSort(nums []int) {
	len := len(nums)
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0 && nums[j] > nums[j+1]; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
}
