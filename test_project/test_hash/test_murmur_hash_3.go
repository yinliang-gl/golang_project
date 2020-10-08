package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

func main() {
	result := uint64(murmur3.Sum64([]byte("fsafdsafdsafadfrqwtryrthnhkuo7i67532412")))
	fmt.Println(result)
}
