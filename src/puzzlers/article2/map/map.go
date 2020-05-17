package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var scene sync.Map
// 	scene.Store("greece", 97)
// 	scene.Store("london", 100)
// 	scene.Store("egypt", 200)

// 	fmt.Println(scene.Load("london"))
// 	scene.Delete("london")
// 	//遍历该map，参数是一个函数，该函数的两个参数是遍历获得的key和value,返回一个bool值，当bool值未false时，遍历结束。
// 	scene.Range(func(k, v interface{}) bool{
// 		fmt.Println("iterate:", k, v)
// 			return true
// 	})
// }

func main() {
	var m sync.Map
	m.Store("lihao", "fe")
	fmt.Println(m.Load("lihao"))
	m.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
