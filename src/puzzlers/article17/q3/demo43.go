package main

import "fmt"

func main() {
	// 示例1。 switch表达式不允许case表达式的子表达式结果值存在相等的情况，针对由字面量直接表示的子表达式
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value3[4] { // 这条语句无法编译通过。
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	// 示例2。
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or26")
	}

	// 示例3。
	//value6 := interface{}(byte(127))
	//switch t := value6.(type) { // 这条语句无法编译通过。
	//case uint8, uint16:  //byte类型是uint8类型的别名类型。因此，byte类型和uint8类型重复
	//	fmt.Println("uint8 or uint16")
	//case byte:
	//	fmt.Printf("byte")
	//default:
	//	fmt.Printf("unsupported type: %T", t)
	//}
}
