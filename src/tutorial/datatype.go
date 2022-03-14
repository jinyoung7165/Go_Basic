package main

func main() {
	var num1 int
	var str1 string
	var num2 int = 7
	var str2 = "Hello, World!" //string 생략하면 알아서 추론
	var i, j, k int = 1, 2, 3
	t := 2 //함수 안에서는 :=도 가능, 함수 밖에서는 var 사용해야함
	const c int = 10
	const (
		num10 = 1
		visa  = 12
	)
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	const ( //iota 지정 시 나머지 두 개는 1부터 순차적으로 증가된 값을 가지게 됨
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)
	println(num1, num2, str1, str2, i, j, k, t) //var는 사용되지 않으면 에러 일으킴
	/*
		문자열 타입
	*/
	const s string = "Hi" //string은 한 번 생성되면 수정될 수 없는 immutable 타입, 플러스 연산 가능
	/*
		'' 안의 것은 raw String Literal, "" 안의 것은 통상적인 string 인용부호
		''안에 엔터 가능, /n 써도 엔터로 인식 x
		""안에 /n 쓰면 엔터로 출력
	*/
	/*
		데이터 타입 변환 : 반드시 명시적 변환!!!!!!!!!!!!!
	*/

	/*
		포인터
		var k int = 10
		var p = &k  //k의 주소를 할당
		println(*p) //p가 가리키는 주소에 있는 실제 내용을 출력
	*/

}
