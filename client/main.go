package main

import (
	"flag"
	"log"
)

var (
	addr = ":7001"
)

var (
	userName = flag.String("name", "", "user name")
)

func init() {
	flag.Parse()
}

func main() {
	if "" == *userName {
		log.Fatal("input your name")
	}
	NewClient(addr).Run()
}
