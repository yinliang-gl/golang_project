package main

import "fmt"

type AAA struct {
	a int
	b int64
}

func test(aaa AAA) {
	fmt.Printf("%d, %d\n", aaa.a, aaa.b)
}

func main() {

}
