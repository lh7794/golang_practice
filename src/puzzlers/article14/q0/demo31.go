package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
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
	_, ok := interface{}(dog).(Pet) //dog类型本身的方法集合只包含了2个方法，即所有的值方法
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet) //*dog类型包含3个方法，值方法和指针方法，所以*dog类型是Pet类型
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	var pet Pet = &dog //*dog才是Pet类型，dog不是pet类型，于demo32.go比较
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

	dog = Dog{"little pig"}

	pet1 := dog //dog类型不是pet接口类型，为什么也可以赋值？？
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet1.Category(), pet1.Name())

	_, ok = interface{}(dog).(Pet) //实现了以上4个方法就属于pet接口
	fmt.Printf("dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*dog implements interface Pet: %v\n", ok)
}
