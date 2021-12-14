package main

import (
	"fmt"
	"os/user"
	"time"
)

// 一番初めに呼ばれる
func init() {
	println("init")
}

//
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sub1()
	now_date()
	osuser()
}

func sub1() {
	fmt.Println("Yahoo")
	fmt.Println("Yahoo", "!!")
}

func now_date() {
	nowtime := time.Now()
	fmt.Println("The time is", nowtime.Format(time.RFC3339))
}

func osuser() {
	fmt.Println(user.Current())
}
