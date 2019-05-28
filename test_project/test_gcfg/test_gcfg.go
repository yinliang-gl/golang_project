package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
)

type GlobalContext struct {
	Server struct {
		Hostname  string
		Host      string
		Socket    string
		MaxConn   int
		PprofHost string
	}

	IndexFile struct {
		Filepath      string
		CheckInterval int
	}

	Banlist struct {
		Filename      string
		CheckInterval int
	}
}

func main() {
	var globalContext GlobalContext
	err := gcfg.ReadFileInto(&globalContext, "test_project/test_gcfg/test.conf")

	if err != nil {
		log.Fatalf("init global conf fail . err[%s]", err.Error())
	}
	fmt.Println(globalContext)
	fmt.Println(globalContext.IndexFile.Filepath)
}
