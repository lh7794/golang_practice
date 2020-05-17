// package main

// import (
// 	"errors"
// 	"fmt"
// )
package main

import "fmt"

func gencalculator() func(x, y int) (int, error) {
	return func(x, y int) (int, error) {//func(x, y int) (int, error)为匿名函数
		return x + y, nil
	}
}
func main() {
	x, y := 12, 23
	add := gencalculator()
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}

// type operate func(x, y int) int
// type calculateFunc func(x int, y int) (int, error)

// func genCalculator(op operate) calculateFunc {
// 	return func(x int, y int) (int, error) {
// 		if op == nil {
// 			return 0, errors.New("invalid operation")
// 		}
// 		return op(x, y), nil
// 	}
// }

// func main() {
// 	x, y := 56, 78
// 	op := func(x, y int) int {
// 		return x + y
// 	}
// 	add := genCalculator(op)
// 	result, err := add(x, y) //result和err为genCalculator函数的返回值
// 	fmt.Printf("The result: %d (error: %v)\n",
// 		result, err)
// }
