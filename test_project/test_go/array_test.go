package test_go

import (
	"fmt"
	"reflect"
	"testing"
)

/**
测试 append
*/
func TestArray001(t *testing.T) {
	fmt.Println("fdsafsd")
	test8_3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := test8_3[2:]

	fmt.Println("old T", test8_3, len(test8_3), cap(test8_3))  // old T [0 1 2 3 4 5 6 7 8 9] 10 10
	fmt.Println("old A", a, len(a), cap(a), reflect.TypeOf(a)) // old A [2 3 4 5 6 7 8 9] 8 8 []int

	// append 数字
	a = append(a, 10) // 注意这个append， 如果这个append被注释掉那么对 a 的改动会影响 test8_3，否则不会影响 test8_3
	a[2] = 100
	fmt.Println("new A", a, len(a), cap(a))                   //new A [2 3 100 5 6 7 8 9 10] 9 16
	fmt.Println("new T", test8_3, len(test8_3), cap(test8_3)) //new T [0 1 2 3 4 5 6 7 8 9] 10 10

	// append 数组
	b := []int{-1, -2, -3}
	a = append(a, b[1:]...)
	fmt.Println("new A after append b", a, len(a), cap(a))
}
