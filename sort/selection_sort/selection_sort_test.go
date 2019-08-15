package algorithm

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}

	fmt.Println("before selection sort", nums)
	SelectionSort(nums)
	fmt.Println("after selection sort", nums)
}
