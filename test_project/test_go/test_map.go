package main

import "fmt"

type MaterialResource struct {
	ID          int32
	Type        int32
	FileName    string
	RemoteURL   string
	FileSize    int64
	ImageWidth  int32
	ImageHeight int32
	AppName     string
	AppPackage  string
	AppLogo     int32
	AppVersion  string
	FileMd5     string
	Duration    uint32
}

func main() {
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
