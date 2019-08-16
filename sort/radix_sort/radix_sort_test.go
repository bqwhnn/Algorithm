package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	nums := []int{321, 1, 10, 60, 743, 127, 577, 2, 24, 45, 66, 75, 90, 170, 802}

	fmt.Println("before radix sort", nums)
	nums = RadixSort(nums)
	fmt.Println("after radix sort", nums)
	fmt.Println("issorted", sort.IntsAreSorted(nums))
}
