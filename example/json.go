package main
import (
	"fmt"
	"encoding/json"
	"regexp"
)
func main() {
	TestCreateSingleItemResponse()
}


//type TestObject struct {
//	Kind string //`json:"kind"`
//	Id   string `json:"id, omitempty"`
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}
type TestObject struct {
	Kind string
	Id   string
	Name  string
	Email string
	Stu  []Stu
}

type Stu struct {
	No int16
	Class string
}

func TestCreateSingleItemResponse() {

	matched, err := regexp.MatchString("zdd", "zddhub")
	fmt.Println(matched, err)
	t := new(TestObject)
	t.Kind = "TestObject"
	t.Id = "f73h5jf8"
	t.Name = "Yuri Gagarin"
	t.Email = "Yuri.Gagarin@Vostok.com"
	s1 := Stu{No:12121}
	s2 := Stu{Class:"fdafafa"}
	t.Stu = append(t.Stu, s1)
	t.Stu = append(t.Stu, s2)
	fmt.Println(t)

	b, err := json.Marshal(t)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b[:]))
}
//Note that you don’t need parentheses around conditions in Go, but that the braces are required.
//注意：你不需要在条件上加括号，但是语句体上是一定要加上的
//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
//Go中没有三元运算表达式，所以得用if条件语句进行替换