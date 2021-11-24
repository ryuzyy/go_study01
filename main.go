package main

import "fmt"

func init() {
	println("init")
}

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
