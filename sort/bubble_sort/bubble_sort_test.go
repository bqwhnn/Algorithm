package algorithm

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}

	fmt.Println("before bubble sort", nums)
	BubbleSort(nums)
	fmt.Println("after bubble sort", nums)
}

func TestBubbleSort2(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}

	fmt.Println("before bubble sort 2", nums)
	BubbleSort2(nums)
	fmt.Println("after bubble sort 2", nums)
}
