package algorithm

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	nums := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 04}

	fmt.Println("before shell sort", nums)
	ShellSort(nums)
	fmt.Println("after shell sort", nums)
}
