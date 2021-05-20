package main

import "fmt"

type Rectangle1 struct {
	width  int
	height int
}

type Rectangle2 struct {
	width, height int
}

func main() {
	var rect1 *Rectangle1   //구조체 포인터 선언
	rect1 = new(Rectangle1) //구조체 포인터에 메모리 할당
	fmt.Println(rect1)

	rect2 := new(Rectangle1) //구조체 포인터 선언과 동시에 메모리 할당
	fmt.Println(rect2)

	var rect3 Rectangle1 = Rectangle1{10, 20} //구조체 인스턴스를 생성하면서 값 초기화
	fmt.Println(rect3)

	rect4 := Rectangle1{45, 62} //var 키워드 생략
	fmt.Println(rect4)

	rect5 := Rectangle1{width: 30, height: 15} //구조체 필드 지정하여 값 초기화
	fmt.Println(rect5)

}
