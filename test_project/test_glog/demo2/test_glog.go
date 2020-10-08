package main

import (
	"flag"
	"github.com/golang/glog"
	"time"
)

func main() {
	defer glog.Flush()
	flag.Lookup("log_dir").Value.Set("log")
	flag.Lookup("log_dir").Value.Set("log")
	flag.Parse()

	for {
		glog.Error("This is error message")
		time.Sleep(1 * time.Second)
	}
}
