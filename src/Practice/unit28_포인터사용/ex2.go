package main

import "fmt"

func main() {
	var num int = 1
	var numPtr *int = &num //참조로 num 변수의 메모리 주소를 구하여 numPtr 포인터 변수에 대입

	fmt.Println(numPtr)
	fmt.Println(&num)
}
