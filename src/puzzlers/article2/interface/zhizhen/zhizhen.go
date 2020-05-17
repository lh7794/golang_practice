package main

import "fmt"

func main() {
	str := new(string) //str为指针类型

	*str = "nanjing"
	fmt.Println(str) //str就是地址
	//	fmt.Println(&str) //str是地址的地址
	fmt.Println(*str) //*str是变量的值
}
