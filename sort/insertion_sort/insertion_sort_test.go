package algorithm

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}

	fmt.Println("before insertion sort", nums)
	InsertionSort(nums)
	fmt.Println("after insertion sort", nums)
}
