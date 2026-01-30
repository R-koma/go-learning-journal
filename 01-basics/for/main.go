package main

import "fmt"

/*
Goに、whileはない
*/

func main() {
	res1 := 0
	for i := 0; i < 100; i++ {
		res1 += i
	}
	fmt.Println(res1)

	// Go 1.22以降の記述方法
	res2 := 0
	for j := range 100 {
		res2 += j
	}
	fmt.Println(res2)
}