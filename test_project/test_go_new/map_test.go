package test_go_new

import (
	"fmt"
	"testing"
)

type MaterialResource struct {
}

func TestMap01(t *testing.T) {
	map_variable := make(map[string]*MaterialResource)
	if map_variable["aaa"] != nil {
		fmt.Println("aaa1 found")
	} else {
		fmt.Println("aaa1 not found")
	}

	if map_variable["aaa"] != nil {
		fmt.Println("aaa2 found")
	} else {
		fmt.Println("aaa2 not found")
	}
}
