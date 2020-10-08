package main

import (
	"fmt"
)

func main() {
	requests := []int{12, 2, 3, 41, 5, 6, 1, 12, 3, 4, 2, 31}
	for _, n := range requests {
		run(n)
	}

	for i:=0; i < 100; i++ {
		run(i)
	}

	for {
		select {}
	}
}

func run(num int) {
	defer func() {
		if err := recover();err != nil {
			fmt.Printf("%s\n", err)
		}
	}()

	fmt.Printf("num:%d\n", num)
	//模拟请求错误
	if num%5 == 0 {
		panic("请求出错")
	}

}
