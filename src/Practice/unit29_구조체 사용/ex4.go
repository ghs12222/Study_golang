package main

import "fmt"

type Rectangle struct {
	width, height int
}

//함수에서 구조체 매개변수 사용
func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func main() {
	rect := Rectangle{30, 30}

	area := rectangleArea(&rect)
	fmt.Println(area)
}
