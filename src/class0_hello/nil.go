package main

import (
	"fmt"
	"unsafe"
)

var abc = 1

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println(unsafe.Sizeof(i))
}
