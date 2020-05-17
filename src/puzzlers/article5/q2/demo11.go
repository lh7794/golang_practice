package main

import "fmt"

var container1 = []string{"zero", "one", "two"}
var container = map[int]string{0: "zero", 1: "lihao", 2: "two"}
var container2 = map[int]string{0: "zero", 1: "lihao", 2: "two"}

func printSlice(x []string) {
	fmt.Printf("container1 is %s.", x[1:])
}
func main() {

	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
	fmt.Printf("The element is %q.\n", container[1])//不同代码块中的重名变量被屏蔽
	fmt.Printf("The element is %q.\n", container2[1])

	printSlice(container1) //参数为container时，会报错
}
