package main

import "fmt"

type Rectangle struct {
	width, height int
}

// 구조체 생성자 패턴
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} //구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func main() {
	rect := NewRectangle(20, 10)

	fmt.Println(rect)

	rect1 := &Rectangle{20, 10}
	fmt.Println(rect1)
}
