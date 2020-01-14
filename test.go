package main

import "fmt"

type AAA struct {
	aaa int32
	bbb int64
	ccc float64
}

func main() {
	var val []AAA
	size := 10

	val = make([]AAA, size, size)

	for index, _ := range val {
		val[index].aaa = int32(index + 1)
	}
	fmt.Println(val)
}
