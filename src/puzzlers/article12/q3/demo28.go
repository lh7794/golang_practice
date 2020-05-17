package main

import "fmt"

func main() {
	// 示例1。
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Println()

	// 示例2。
	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
	fmt.Printf("The modified slice: %v\n", slice2) //切片原值和修改后值都会改变
	fmt.Printf("The original slice: %v\n", slice1)
	fmt.Println()

	// 示例3。
	complexArray1 := [3][]string{ //切片数组
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)

	// 示例4
	ca1 := [][3]string{
		[3]string{"1", "2", "3"},
		[3]string{"4", "5", "6"},
		[3]string{"7", "8", "9"},
	}
	fmt.Printf("The ca array: %v\n", ca1)
	ca2 := modiftca(ca1)
	fmt.Printf("The modified complex array: %v\n", ca2)
	fmt.Printf("The original complex array: %v\n", ca1)
}

// 示例1。
func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

// 示例2。
func modifySlice(a []string) []string {
	a[1] = "i"
	return a
}

// 示例3。
func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"                  //  指针未改变
	a[2] = []string{"o", "p", "q"} // 指针发生改变 修改的是复制的数据的副本
	return a
}

// 示例4
func modiftca(a [][3]string) [][3]string {
	a[1][1] = "s" //指针未改变  外层是指针 值全部发生改变
	a[2] = [3]string{"a", "b", "c"}// 
	return a
}
