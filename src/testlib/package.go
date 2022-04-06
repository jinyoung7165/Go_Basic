package testlib //사용자정의 package(폴더 이름과 같음)
//패키지 폴더 안에 여러 파일들이 있을 때도 동일하게 testlib 패키지명을 사용
import "fmt"

var pop map[string]string

func init() { //패키지 로드 시 map 초기화
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
    pop["Alicia Keys"] = "Fallin'"
    pop["John Legend"] = "All of Me"
}

func GetMusic(singer string) string { //대문자로 시작-> 패키지 import 시 외부에서 호출 가능
	return pop[singer]
}

func getKeys() { //내부에서만 호출 가능
	for _, kv := range pop {
		fmt.Println(kv)
	}
}