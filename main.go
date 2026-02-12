package main

import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	numCPU := runtime.NumCPU()
	fmt.Println("Mallocs: ", rtm.Mallocs)
	fmt.Println("HeapAlloc: ", rtm.HeapAlloc)
	fmt.Println("numCPU: ", numCPU)	
}
