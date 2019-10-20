package main

import "fmt"

/**
测试三个点的用法
*/
func test_basic_grammar_001() {
	var strss = []string{
		"qwr",
		"234",
		"yui",
	}
	var strss2 = []string{
		"qqq",
		"aaa",
		"zzz",
		"zzz",
	}
	strss = append(strss, strss2...) //strss2的元素被打散一个个append进strss
	fmt.Println(strss)
}
func test_basic_grammar_002() {
	var strss = []string{
		"qwr",
		"234",
		"yui",
		"cvbc",
	}

	iterater_fun := func(args ...string) { //可以接受任意个string参数
		for _, v := range args {
			fmt.Println(v)
		}
	}

	iterater_fun(strss...) //切片被打散传入
}

func main() {
	// test_basic_grammar_001();
	test_basic_grammar_002()
}
