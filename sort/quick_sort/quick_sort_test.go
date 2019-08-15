package algorithm

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 04}

	fmt.Println("befort quick sort 1", nums)
	QuickSort(nums)
	fmt.Println("after quick sort 1", nums)
}

func TestQuickSort2(t *testing.T) {
	nums := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 04}

	fmt.Println("befort quick sort 2", nums)
	QuickSort2(nums)
	fmt.Println("after quick sort 2", nums)
}
