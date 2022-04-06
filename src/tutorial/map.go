package main

import "fmt"

func main() { //map: k, v 해쉬테이블
	var m map[int]string
	m = make(map[int]string)
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	delete(m, 777)

	str := m[134] //값만 할당(복사)
	println(str)

	m[134] = "Hey"
	println(str)

	
	tickers := map[string]string{
		"Goog": "GOOGLE",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	//map 키 체크
	val, exists := tickers["MSFT"] //값, t/f
	println(val, "hey!!!", exists)
	if !exists {
		println("No MSFT tickers")
	}

	//열거
	for key, val := range tickers {
        fmt.Println(key, val)
    }



	
}