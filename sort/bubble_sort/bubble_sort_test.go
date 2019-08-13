package algorithm

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}

	fmt.Println("before bubble sort", nums)
	bubbleSort(nums)
	fmt.Println("after bubble sort", nums)
}
