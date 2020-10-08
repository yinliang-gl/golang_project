package test_go

import (
	"fmt"
	"testing"
)

/**
make的第三个参数的作用，
make函数第一个参数是类型，第二个参数是分配的空间，第三个参数是预留分配空间，
例如a:=make([]int, 5, 10)， len(a)输出结果是5，cap(a)输出结果是10，然后对a[4]进行赋值发现是可以的，
但对a[5]进行赋值发现报错了，原因是预留的空间需要重新切片才可以使用
*/
func TestMake01(t *testing.T) {
	a := make([]int, 10, 20)
	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	b := a[:cap(a)]
	fmt.Println(b)
}

/**
map的make
*/
func TestMake02(t *testing.T) {
	goroutineMap := make(map[string]int)
	goroutineMap["goroutine"] = 10
	goroutineMap["goroutine"] = 11
	fmt.Println(goroutineMap)
}
