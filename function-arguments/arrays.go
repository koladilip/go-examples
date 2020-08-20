package main

import (
	"fmt"
)

func incArrays(nums [10]int) [10]int {
	for i := 0; i < len(nums); i++ {
		nums[i]++
	}
	return nums
}

func main() {
	nums := [10]int{}
	fmt.Println("before", nums)
	fmt.Println("returned", incArrays(nums))
	fmt.Println("after", nums)
}
