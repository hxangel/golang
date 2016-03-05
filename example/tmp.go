package main

import (
	"strings"
	"fmt"
	"strconv"
	"math/rand"
//	"time"
	"time"
)



func main() {
	id := GetExtensionId()
	fmt.Println(id)
}
func GetExtensionId() string {
	s4 := func() string {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		 str := ToHex(r.Intn(rand.Int()), 2)
		return str[1:5]

	}
	s := s4() + s4() + "-" + s4() + "-" + s4() + "-" + s4() + "-" + s4() + s4() + s4()
	return s;
}
//字符串反转
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
//遍历msg的字符,转换成相应with位16进制,然后连接在一起
func Encrypt(msg string, with int, reverse bool) string {
	var ch_arr []string
	for _, ch := range msg {
		ch_hex := ToHex((int(ch) + 88), with)
		ch_arr = append(ch_arr, ch_hex)
	}
	ret := strings.Join(ch_arr, "")
	if reverse {
		return Reverse(ret)
	}
	return ret
}
//当前字符+88,变成对应位数的16进制
func ToHex(ch int, width int) string {
	return fmt.Sprintf("%0" + strconv.Itoa(width) + "x", ch)
}