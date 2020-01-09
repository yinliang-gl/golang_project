package main

import "fmt"

type AAA struct {
	a        int
	b        int64
	test_map map[int32]int32
}

func test(aaa AAA) {
	fmt.Printf("%d, %d\n", aaa.a, aaa.b)
}

func Run(aaa bool) (isFiltered bool) {
	if !aaa {
		isFiltered = true
	}
	return
}

func main() {
	isFiltered := false

	isFiltered = Run(true)
	fmt.Println(isFiltered)

	isFiltered = Run(false)
	fmt.Println(isFiltered)
}
