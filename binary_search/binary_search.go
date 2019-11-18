package binarysearch

// BinarySearch func
func BinarySearch(a []int, target int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + (high-low)/2
		if a[mid] > target {
			high = mid - 1
		} else if a[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
