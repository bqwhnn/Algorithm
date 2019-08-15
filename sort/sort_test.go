package sort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	// fmt.Println(generateRandomArray(100))
	fmt.Println(generateNearlyOrderedArray(10))
}

// 生成随机数列
func generateRandomArray(n int) []int {
	a := make([]int, n)
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}
	return a
}

// 生成基本有序的数列
func generateNearlyOrderedArray(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < n/3; i++ {
		x := rand.Intn(n - 1)
		y := rand.Intn(n - 1)
		fmt.Println(x, y)
		a[x], a[y] = a[y], a[x]
	}
	return a
}

// 判断数列是否有序
func isSorted(a []int) bool {
	return sort.IntsAreSorted(a)
}
