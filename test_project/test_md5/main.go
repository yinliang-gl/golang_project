package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	println("fasfasd")

	androidId := "devId.ID"
	androidMd5 := fmt.Sprintf("%x", md5.Sum([]byte(androidId)))

	println(androidMd5)
}
