package main

import (
	"fmt"
)

func main() {
	typeSpecify("a")
}

func typeSpecify(o interface{}) {
	fmt.Printf("%T", o)
}
