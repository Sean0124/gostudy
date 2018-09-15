package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	test()
}

func newInt01() *int {
	return new(int)
}

func newInt02() *int {
	var dummy int
	return &dummy
}

func test() {
	p := new(int)
	q := new(int)
	fmt.Println(p == q)
	fmt.Println(*p == *q )
}