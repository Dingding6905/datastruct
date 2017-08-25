package main

import (
	"fmt"
)

// 折半查找
func BinargSearch(a []int, n int, key int) int {
	mid := 0
	low := 0
	high := n - 1
	for low <= high {
		mid = (low + high) / 2
		if key > a[mid] {
			fmt.Println("---mid=", mid)
			low = mid + 1
		} else if key < a[mid] {
			fmt.Println("===mid=", mid)
			high = mid - 1
		} else {
			fmt.Println("xxx mid=", mid)
			return mid
		}
	}
	return 0
}

func main() {
	i := make([]int, 10)
	copy(i, []int{1, 2, 5, 7, 9, 11, 15, 176})
	n := 10
	fmt.Println(BinargSearch(i, n, 15), n, i)
}
