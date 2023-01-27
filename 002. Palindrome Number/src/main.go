// package main

// import (
// 	"fmt"
// )

func isPalindrome(x int) bool {
	result := true
	arr := make([]int, 0)
	if x < 0 {
		return false
	}
	for x > 0 {
		i := x % 10
		arr = append(arr, i)
		x = x / 10
	}
	l := len(arr)
	var n int = l / 2
	for i := 0; n > 0; i++ {
		if arr[i] != arr[l-1] {
			return false
		} else {
			n--
			l--
		}
	}
	return result
}

// func main() {
// 	r := isPalindrome(123321)
// 	fmt.Println(r)
// }
