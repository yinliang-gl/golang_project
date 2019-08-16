package main

import (
	"fmt"
	"runtime"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func PrintHello() {
	fmt.Println("Hello world")
}

func GetChars(s string) {
	for _, c := range s {
		fmt.Printf("%c ,time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func GetDigit(s []int) {
	for _, d := range s {
		fmt.Printf("%d , time %v\n", d, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	fmt.Println("main start ,time ", time.Since(start))

	go GetChars("abcde")
	go GetDigit([]int{1, 2, 3, 4, 5})

	time.Sleep(1 * time.Second)

	fmt.Println("\nmain end, time  ", time.Since(start))
}
