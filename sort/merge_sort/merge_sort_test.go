package algorithm

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 04}

	fmt.Println("before merge sort", nums)
	MergeSort(nums)
	fmt.Println("after merge sort", nums)
}
