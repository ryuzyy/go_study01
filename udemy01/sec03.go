package main

import "fmt"

// var (
// 	i    int     = 1
// 	f64  float64 = 1.2
// 	s    string  = "test"
// 	t, f bool    = true, false
// )

// func foo() {
// 	xi := 1
// 	xf64 := 1.2
// 	xs := "test"
// 	xt, xf := true, false
// 	fmt.Println(xi, xf64, xs, xt, xf)
// 	fmt.Printf("%T", xi)
// }

// func sd_const() {
// 	const Pi = 3.14

// 	const (
// 		Username = "test_user"
// 		Password = "test_pass"
// 	)
// 	fmt.Println(Pi, Username, Password)
// }

// func sd_make() {
// 	var c []int
// 	c = make([]int, 5)
// 	// c = make([]int, 0, 5)
// 	for i := 0; i < 5; i++ {
// 		c = append(c, i)
// 		fmt.Println(c)
// 	}
// 	fmt.Println(c)
// }

func sd_map() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	println(v, ok)

	v2, ok2 := m["nothing"]
	println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 50000
	fmt.Println(m2)

	// var m3 map[string]int
	// m3["pc"] = 50000
	// fmt.Println(m3)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}

func main() {
	// sd_const()
	// sd_make()
	sd_map()
}
