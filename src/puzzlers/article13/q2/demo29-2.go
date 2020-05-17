package main

import (
	"fmt"
)

type Vertex struct {
	X, Y string
}

func (v Vertex) String() string {
	return fmt.Sprintf("%s%s", v.X, v.Y)

}

func main() {
	v := Vertex{Y: "4"}
	fmt.Println(v)
}
