package algorithm

// BubbleSort is basic bubble sort
func BubbleSort(nums []int) {
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

// BubbleSort2 is advanced bubble sort
func BubbleSort2(nums []int) {
	len := len(nums)

	flag := 0
	for flag < len {
		k := flag
		flag = len
		for j := len - 1; j > k; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				flag = j
			}
		}
	}
}
