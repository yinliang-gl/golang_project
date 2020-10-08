package test_go

import (
	"fmt"
	"testing"
)

type Father interface {
	Run()
}

type Child1 struct {
}

func (p *Child1) Run() {
	fmt.Println("Child1 running")
}

type Child2 struct {
}

func (p *Child2) Run() {
	fmt.Println("Child2 running")
}

func Get(index int) interface{} {
	if index%2 == 0 {
		return Child1{}
	} else {
		return Child2{}
	}
}

func TestTypeConversion01(t *testing.T) {
	for index := 0; index < 5; index++ {
		bbb := Get(index)
		_, res1 := bbb.(Child1)
		_, res2 := bbb.(Child2)
		if res1 == true {
			fmt.Printf("element of index[%d] is Child1\n", index)
		} else if res2 == true {
			fmt.Printf("element of index[%d] is Child2\n", index)
		} else {
			fmt.Println("not child1 and not child2")
		}
	}
}
