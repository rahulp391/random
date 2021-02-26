package main

import (
	"fmt"
	"sort"
)

func median(a ...int) float32 {
	sort.Ints(a)
	var s float32
	l := len(a)
	if l == 0 {
		return 0
	}
	if l%2 == 0 {
		s = float32(a[int(l/2)] + a[int(l/2)-1])
		s /= 2
	} else {
		s = float32(a[int(l/2)])
	}
	return s
}
func main() {
	a := []int{4, 5, 6, 7, 2, 1}
	fmt.Println(median(a...))

}
