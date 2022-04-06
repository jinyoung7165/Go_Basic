package main

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	//배열 초기화
	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3} //배열 크기 자동

	//다차원 배열
	var multiArray [3][4][5]int //정의
	multiArray[0][1][2] = 10    //사용

	var arr = [2][3]int{ //다차원 배열 초기화
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가!!!!!!
	}
	println(arr[1][2])
}