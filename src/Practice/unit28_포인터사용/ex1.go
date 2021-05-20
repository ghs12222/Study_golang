package main

import "fmt"

func main() {

	var numPtr *int //포인터형 변수 선언 시 nil로 초기화
	fmt.Println(numPtr)

	var numPtr2 *int = new(int)
	fmt.Println(numPtr2) //메모리 주소.

	var numPtr3 *int = new(int)
	*numPtr3 = 1 //역참조로 포인터형 변수에 값을 대입
	fmt.Println(*numPtr3)
}
