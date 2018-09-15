package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	//定义一个map,key为int型,value值为string型,map内部无序
	fmt.Println(m1)
	//打印m1
	m2 := make(map[string]int)
	//make一个,key为string,value为int的map

	for k, v := range m1 {
	//for循环,将m1的key和value通过range迭代赋值给k和v
		m2[v] = k
	//用v作为key值,用k作为value,赋值给m2
	}
	fmt.Println(m2)
	//打印m2
}
