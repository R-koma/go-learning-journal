package main

import "fmt"

func main(){
	defer fmt.Println("after")
	fmt.Println("before")
}
