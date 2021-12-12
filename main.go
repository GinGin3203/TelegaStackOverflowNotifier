package main

import "fmt"
import "runtime"

func T() {}
func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
}
