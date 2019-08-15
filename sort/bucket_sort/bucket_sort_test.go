package algorithm

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	nums := []int{}

	fmt.Println("before bucket sort", nums)
	BucketSort(nums)
	fmt.Println("after bucket sort", nums)
}
