package main

import (
	"fmt"
)

func incSlice(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i]++
	}
	return nums
}

func main() {
	nums := make([]int, 10)
	fmt.Println("before", nums)
	fmt.Println("returned", incSlice(nums))
	fmt.Println("after", nums)
}
