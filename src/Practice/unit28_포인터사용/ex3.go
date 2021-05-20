package main

import "fmt"

func hello1(n int) {
	n = 2
}

func hello2(n *int) {
	*n = 2
}

func main() {
	var n int = 1
	hello1(n)
	fmt.Println(n)
	hello2(&n)
	fmt.Println(n)
}
