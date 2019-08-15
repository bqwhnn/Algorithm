package algorithm

// ShellSort is shell sort
func ShellSort(nums []int) {
	len := len(nums)
	for gap := len / 2; gap > 0; gap /= 2 {
		for i := gap; i < len; i++ {
			for j := i - gap; j >= 0 && nums[j] > nums[j+gap]; j -= gap {
				nums[j], nums[j+gap] = nums[j+gap], nums[j]
			}
		}
	}
}
