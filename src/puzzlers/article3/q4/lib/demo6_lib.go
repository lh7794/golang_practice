package lib

import (
	"os"

	in "geektime/Golang_Puzzlers/src/puzzlers/article3/q4/lib/internal"
)

//hu
func Hello(name string) {

	in.Hello(os.Stdout, name)
}
