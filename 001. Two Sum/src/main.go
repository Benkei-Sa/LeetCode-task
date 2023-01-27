// package main

// import (
// 	"fmt"
// )

func twoSum(nums []int, target int) []int {
	i_sum := []int{}
	nums_l := len(nums)
	for i := 0; i < nums_l; i++ {
		for n := (i + 1); n < nums_l; n++ {
			if (nums[i] + nums[n]) == target {
				i_sum = append(i_sum, i, n)
				return i_sum
			}
		}
	}
	return nil
}

// func main() {
// 	nums := []int{3, 2, 4}
// 	var target = 6
// 	fmt.Println(twoSum(nums, target))
// }