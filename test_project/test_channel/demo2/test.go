package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string
}

/*
测试channel传输复杂的Struct数据
这里可以看到可以通过channel传输自定义的Person对象，同时一端修改了数据，不影响另一端的数据，也就是说通过channel传递后的数据是独立的。
*/
func testTranslateStruct() {
	personChan := make(chan Person, 1)

	person := Person{"xiaoming", 10, Addr{"shenzhen", "longgang"}}
	personChan <- person

	person.Address = Addr{"guangzhou", "huadu"}
	fmt.Printf("src person : %+v \n", person)

	newPerson := <-personChan
	fmt.Printf("new person : %+v \n", newPerson)
}

func main() {
	testTranslateStruct()
}
