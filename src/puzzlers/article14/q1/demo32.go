package main

import (
	"fmt"
)

type Pet interface { //Pet类型有两个值方法，所以dog类型是pet接口类型
	//	SetName(name string)
	Name() string
	Category() string
}

type Dog struct { //dog类型本身的方法集合只包含了2个方法，值方法。*dog类型包含了3个方法，值方法和指针方法
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	var pet Pet = dog //dog类型是pet类型的接口实现
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name()) //pet变量的字段name的值依然是"little pig"，值未发生改变 因为传递的是该值的副本
	fmt.Println()

	// 示例2。
	dog1 := Dog{"little pig"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	// 示例3。
	dog = Dog{"little pig"} //dog就是一个普通的变量，所以可以去地址哦！
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog // pet.name的值发生了改变 ？？？？
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}
