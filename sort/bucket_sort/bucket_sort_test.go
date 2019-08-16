package algorithm

import (
	"fmt"
	"testing"
	"sort"
)

func TestBucketSort(t *testing.T) {
	nums := []int{29, 25, 3, 49, 9, 10, 5, 4, 9, 37, 21, 43, 69, 58, 1}

	fmt.Println("before bucket sort", nums)
	nums = BucketSort(nums)
	fmt.Println("after bucket sort", nums)
	fmt.Println("issorted", sort.IntsAreSorted(nums))
}
