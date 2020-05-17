package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Name() {

	fmt.Println("dog.name:", dog.name)
}

func newDog(name string) *Dog { //方法二：创建结构体实例
	return &Dog{name}
}

func main() {
	dog1 := new(Dog) //方法一：创建结构体实例 name为空
	dog1.Name()

	dog2 := newDog("zzz") //方法二：创建结构体实例
	dog2.Name()
}
