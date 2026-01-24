package main

import (
	"fmt"
)

func average(x, y int) bool{
	if a := (x + y) / 2; a <= 170 {
		return true
	}
	return false
}

func main() {
	fmt.Println(average(180, 176))
}