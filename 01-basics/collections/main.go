package main

import "fmt"

func main() {
	// arr
	var x = [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println((x[0]))
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// slice
	var y = []int{6, 7, 8, 9, 10}
	fmt.Println(y)
	fmt.Println(y[0])
	fmt.Println(len(y))
	fmt.Println(cap(y))

	y = append(y, 11, 12)
	fmt.Println(y)
	// make
	// len = 6, cap = 12
	z := make([]int, 6, 12)
	fmt.Println(z)

	// Map
	// map[<キーの型>]<値の型>
	// マップリテラル
	numMap := map[string]int{}
	fmt.Println(numMap)
	numMap["a"] = 1
	fmt.Println(numMap)
}