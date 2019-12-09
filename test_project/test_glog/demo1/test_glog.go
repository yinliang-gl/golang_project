package main

import (
	"flag"
	"github.com/golang/glog"
	"time"
)

func main() {
	//flag.Set("log_dir", "log")

	flag.Parse()
	defer glog.Flush()

	for {
		glog.Error("This is error message")
		time.Sleep(1 * time.Second)
	}
}
