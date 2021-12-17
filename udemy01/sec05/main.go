package main

import "fmt"

func sub01() {
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2)
	// *p2++
	// fmt.Println(p2)
}
func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	var p *int = &n

	fmt.Println(p)

	fmt.Println(&p)
	sub01()
}
