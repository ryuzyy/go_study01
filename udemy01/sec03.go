package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T", xi)
}

func sd_const() {
	const Pi = 3.14

	const (
		Username = "test_user"
		Password = "test_pass"
	)
	fmt.Println(Pi, Username, Password)
}

func sd_make() {
	var c []int
	c = make([]int, 5)
	// c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}

func main() {
	sd_const()
	sd_make()
}
