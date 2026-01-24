package main

import "fmt"


var x int = 30

func main() {
	/*
	・型の省略ができる
	・関数内での利用のみでパッケージレベルで変数宣言できない。
	・既存の変数に値を代入できる(varにはできない)
	*/
	y := 30
	fmt.Println(x)
	fmt.Println(y)
}
