package algorithm

func bubbleSort(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}
