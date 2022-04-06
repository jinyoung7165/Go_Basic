package main

import "fmt"

func main() {
	var a[]int //슬라이스 변수 선언 -> 동적으로 크기 변경 가능
	a = []int{1,2,3} //슬라이스에 리터럴 값 지정
	a[1] = 10
	fmt.Println(a) //[1,10,3] 출력

	ss := []int{0, 1, 2, 3, 4, 5}
    ss = ss[2:5]   //부분 슬라이스
    fmt.Println(ss) //2,3,4 출력


	s := []int{0, 1}
	s = append(s, 2)
	s = append(s, 3, 4, 5)
	fmt.Println(s)

	sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}
 
    sliceA = append(sliceA, sliceB...)//슬라이스 병합
    //sliceA = append(sliceA, 4, 5, 6)
 
    fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
}