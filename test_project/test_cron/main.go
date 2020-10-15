package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running i = ", i)
	})
	c.Start()

	time.Sleep(10 * time.Second)

	log.Println("new function")

	j:=0
	c.AddFunc(spec, func() {
		j++
		log.Println("cron running j = ", j)
	})

	select{}
}