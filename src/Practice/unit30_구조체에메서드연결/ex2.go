package main

import "fmt"

type Rectangle struct {
	width, height int
}

//포인터 리시버는 복제 x 레퍼런스가 넘긴다
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

//value 리시버는 값이 복제.
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func main() {
	rect1 := Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2)
}
