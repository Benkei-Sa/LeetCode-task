// package main

// import (
// 	"fmt"
// 	"strings"
// )

func romanToInt(s string) int {
	var result int
	arr_arb := strings.Split(s, "")
	l := len(arr_arb) - 1
	for i := 0; i <= l; i++ {
		switch arr_arb[i] {
		case "M":
			result += 1000
		case "D":
			result += 500
		case "C":
			if i < l && arr_arb[i+1] == "D" {
				result += 400
				i++
			} else if i < l && arr_arb[i+1] == "M" {
				result += 900
				i++
			} else {
				result += 100
			}
		case "L":
			result += 50
		case "X":
			if i < l && arr_arb[i+1] == "L" {
				result += 40
				i++
			} else if i < l && arr_arb[i+1] == "C" {
				result += 90
				i++
			} else {
				result += 10
			}

		case "V":
			result += 5
		case "I":
			if i < l && arr_arb[i+1] == "V" {
				result += 4
				i++
			} else if i < l && arr_arb[i+1] == "X" {
				result += 9
				i++
			} else {
				result += 1
			}
		}
	}
	return result
}

// func main() {
// 	s := "III"
// 	fmt.Println(romanToInt(s))
// }
