package algorithm

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{5, 2, 7, 3, 6, 9, 1, 4, 8}

	fmt.Println("before heap sort", nums)
	HeapSort(nums)
	fmt.Println("after heap sort", nums)
}
