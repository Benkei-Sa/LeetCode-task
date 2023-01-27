package main

import (
	"fmt"
	"strings"
)

// func romanToInt(s string) int {
// 	arr_arb := strings.Split(s, "")
// }

func main() {
	var s string
	fmt.Scanln(&s)
	arr_arb := strings.Split(s, "")
	fmt.Println(arr_arb)
}
