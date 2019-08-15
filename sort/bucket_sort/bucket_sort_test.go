package algorithm

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	nums := []int{29, 25, 3, 49, 9, 37, 21, 43, 69, 58, 1}

	fmt.Println("before bucket sort", nums)
	BucketSort(nums)
	fmt.Println("after bucket sort", nums)
}
