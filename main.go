package main

import "fmt"

// 一番初めに呼ばれる
func init() {
	println("init")
}

// ここから全て始まる
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sub1()
}

func sub1() {
	fmt.Println("Yahoo")
}
