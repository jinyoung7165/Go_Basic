//struct는 필드만을 가지며, method는 따로 정의됨
//메소드는 func키워드와 함수명 사이에 어떤 struct를 위한 메소드인지 표시
package main

type Rect struct {
	width, height int
}
//receiver 어떤 struct인지 적어줌
func (r Rect) area() int { //value receiver
	return r.width * r.height
}

func (r *Rect) area2() int { //ptr receiver
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()//메서드 호출
	area2 := rect.area2()//ptr메서드 호출
	println(area, area2, rect.area())
}