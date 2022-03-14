package main

func main() {
	//for문 초기값;조건식;증감식을 감싸는 괄호()쓰면 에러
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
	//while문 대신 for문에 조건식만 붙여 쓴다
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)
	/* 무한루프
		for {
	        println("Infinite loop")
	    }
	*/

	//for range문 - 컬렉션으로부터 한 요소씩 빼냄
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
	//break, continue, goto 가능
}