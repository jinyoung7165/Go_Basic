package main

/*
pass by value
func main() {
	msg := "Hello"
    say(msg)
}

func say(msg string) {
    println(msg)
}
*/

/*
pass by reference
func main(){
	msg := "Hello"
	say(&msg)
	println(msg) //변경된 메시지 출력
}
func say(msg *string){
	println(*msg)
	*msg = "Changed"
}
*/
/*
	리턴 값이 존재하면 매개변수 괄호 뒤에 리턴값들의 타입을 적는다
*/

/*
func main() {
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)
}

func sum(nums ...int) (int, int) { //두 개 리턴
    s := 0      // 합계
    count := 0  // 요소 갯수
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}
*/

/*
익명함수 변수에 함수 할당 가능
func main() {
    sum := func(n ...int) int { //익명함수 정의
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }

    result := sum(1, 2, 3, 4, 5) //익명함수 호출
    println(result)
}
*/

/*
클로저 - 함수가 외부의 변수를 사용
*/

func nextValue() func() int { //익명함수 func() int 를 리턴하는 newvalue함수 정의
    i := 0
    return func() int { //익명함수이면서 클로저 함수 리턴 (함수 바깥의 i 참조)
        i++
        return i
    }
}
 
func main() {
    next := nextValue() //클로저 함수를 할당
 
    println(next())  // 1
    println(next())  // 2
    println(next())  // 3 i가 계속 증가함
	//next가 변수i를 내부에 유지하고 있음
	
    anotherNext := nextValue() //다시 할당
    println(anotherNext()) // 1 다시 시작
    println(anotherNext()) // 2
}