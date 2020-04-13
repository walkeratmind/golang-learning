package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go is going on:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s\n", os)
	}

}
