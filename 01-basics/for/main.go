package main

import "fmt"

/*
Goに、whileはない
*/

func main() {
	res := 0
	for i := 0; i <= 100; i++ {
		res += i
	}
	fmt.Println(res)
}