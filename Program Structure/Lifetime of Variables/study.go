package main

var global *int

func main() {

}

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
