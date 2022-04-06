package main

import "fmt"

type person struct {
	name string
	age int
}
//struct 개체를 다른 함수의 파라미터로 넘긴다면 객체를 복사해서 전달. 값 변하게 하고 싶으면 struct의 포인터를 전달해야함
//즉, new로 생성
func main() {
	//person객체 생성
	p := person{}
	//필드 값 설정
	p.name = "Lee"
	p.age = 10
	fmt.Println(p)

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p1, p2)

	ptr := new(person) //모든 필드를 zero로 초기화, 객체의 포인터(*person) 리턴
	ptr.name = "Lee" //p가 포인터라도 .을 사용
}

type dict struct {
	data map[int]string
}
//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}
/*
func main {
	dic := newDict() //생성자 호출
	dic.data[1] = "A"
}
*/