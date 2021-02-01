package main

/*
#include <stdlib.h>
*/
import (
	"C"
	"fmt"
)

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	fmt.Println(Random())
}
