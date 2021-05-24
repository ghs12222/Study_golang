package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "hello, world"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i)
		}()
	}

	fmt.Scanln()
}
