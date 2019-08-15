package algorithm

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	nums := []int{321, 1, 10, 60, 743, 127, 577}

	fmt.Println("before radix sort", nums)
	RadixSort(nums)
	fmt.Println("after radix sort", nums)
}
