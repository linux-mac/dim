package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var (
	addr = ":7001"
)

func main() {
	log.Println("running on:", addr)
	err := NewServer(addr).Listen()
	if nil != err {
		log.Fatal(err)
	}
}
