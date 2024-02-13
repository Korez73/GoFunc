package main

import (
	"fmt"
	"os"

	Lib "github.com/Korez73/gofunc/lib"
)

func main() {
	//Lib.Demo()
	//Lib.SliceBlock()
	//Lib.SliceShareStorage()
	//Lib.OuterLabel()

	os.Args = append(os.Args, "test.txt")
	lenargs := len(os.Args)
	fmt.Println(lenargs)
	Lib.ReadFile()
}
