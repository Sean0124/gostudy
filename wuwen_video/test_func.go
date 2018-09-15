package main

import (
	"fmt"
)

func main() {
	var fs =[4]func(){}
	//定义fs类型为函数数组

	for i := 0; i < 4; i++ {
		//for循环
		defer fmt.Println("defer i =", i)
		//按照函数体执行的相反顺序执行,i为值拷贝
		defer func() {fmt.Println("defer_closure i = ", i)}()
		//i为指针拷贝,定义一个匿名函数,形成闭包
		fs[i] = func() {fmt.Println("clusure i = ", i)}
		//i为指针拷贝,定义一个匿名函数,形成闭包,并赋给fs
	}
	for _, f := range fs {
		//迭代输出fs
		f()
	}
}