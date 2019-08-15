package algorithm

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	nums := []int{5, 3, 4, 7, 2, 4, 3, 4, 7}

	fmt.Println("before counting sort", nums)
	fmt.Println("after counting sort", countingSort(nums))
}
