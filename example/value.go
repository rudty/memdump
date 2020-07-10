package main

import (
	"os"
	"time"

	"github.com/rudty/memdump"
)

var value int32

func addValue() {
	for {
		value++
		time.Sleep(100 * time.Millisecond)
	}
}
func subValue() {
	for {
		value--
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go subValue()
	go addValue()
	time.Sleep(1 * time.Second)
	f, _ := os.Create("hello.dmp")
	memdump.WriteFullDump(f)
}
