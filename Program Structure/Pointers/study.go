package main

import (
	"fmt"
)

/*
func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
}


func main(){
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
}


var p = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Println(f() == f())
}
*/

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}