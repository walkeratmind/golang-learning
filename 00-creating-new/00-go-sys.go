package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os: ", runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.CPUProfile())
	// fmt.Println(runtime.MemProfile())
	fmt.Println(runtime.Version())
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOROOT())
}
