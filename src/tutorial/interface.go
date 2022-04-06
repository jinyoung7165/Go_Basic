//구조체는 필드들의 집합
//인터페이스는 메서드의 집합
//interface를 구현하려면 인터페이스가 가진 모든 메소드들을 구현하면 됨
package main

import "math"

//Rect 정의
type Rect struct {
    width, height float64
}
 
//Circle 정의
type Circle struct {
    radius float64
}

type Shape interface { //interface정의
	area() float64
	perimeter() float64
}

//Rect타입에 대한 Shape인터페이스 구현
func (r Rect) area() float64 {return r.width * r.height}
func (r Rect) perimeter() float64 {
	return 2 * ( r.width + r.height )
}

//Circle 타입에 대한 Shape 인터페이스 구현 
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 { 
    return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}
func showArea(shapes ...Shape){ //인터페이스 파라미터
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		println(a)
	}
}

