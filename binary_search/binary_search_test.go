package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 5, 6, 8, 10, 15, 21}
	fmt.Println("new array:", a)
	fmt.Println("search 6:", BinarySearch(a, 6))
	fmt.Println("search 20", BinarySearch(a, 20))
}
