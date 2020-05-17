package main

import (
	"flag"
	"fmt"
)

//var name string

//func init() {
//}

func main() {
	flag.Parse()
	var name = flag.String("name", "everyone", "The greeting object.")

	fmt.Printf("Hello, %s!\n", *name)
}
