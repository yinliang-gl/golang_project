package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	writeToFile2("aaa", "aaa.log")
}

func writeToFile(msg string, path string) {
	f, err := os.OpenFile(path, os.O_WRONLY&os.O_CREATE, 0666)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
	f.Close()
}

// 会覆盖文件
func writeToFile2(msg string, path string) {
	if err := ioutil.WriteFile(path, []byte("hello"), 777); err != nil {
		os.Exit(111)
		log.Println(err.Error())
	}
}
