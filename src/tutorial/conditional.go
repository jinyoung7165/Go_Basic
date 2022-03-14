//조건식에 1,0 불가 (c, python은 가능)
//if와 { 가 같은 줄에 있어야함
//if-else if-else
package main

func main() {
	i,max := 1,4
	if val := i * 2; val < max { //조건식 이전에 문장 실행 가능. val은 이 scope에서만 사용 가능
		println(val)
	}
	var name string
    var category = 1
 
    switch category { //break안 적어도 다른 언어들처럼 다음 case로 가지 않는다. 밑으로 가게 하려면 fallthrough적으면 됨
    case 1:
        name = "Paper Book"
    case 2:
        name = "eBook"
    case 3, 4:
        name = "Blog"
    default:
        name = "Other"
    }
    println(name)
	/*
     
    // Expression을 사용한 경우
    switch x := category << 2; x - 1 {
        //...
    } 
	switch v.(type) {
		case int:
			println("int")
		case bool:
			println("bool")
		case string:
			println("string")
		default:
			println("unknown")
	}   
	*/
}