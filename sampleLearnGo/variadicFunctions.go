//go:build ignore

package main

import (
	"fmt"
)

func main() {

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	nums := []int{10, 11, 12, 13, 14, 15}
	sum(nums...)

	fmt.Println(fact(7))
}

// 可変長引数
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 再帰関数（Recursion)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
