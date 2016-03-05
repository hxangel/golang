package main
import (
	"fmt"
	"strings"
	"strconv"
	"net/url"
)
func main() {
	pw := Cover("sdfadfadfasfdda")
	println(pw)

}

func Cover(item_url string) string {
	url_query := url.QueryEscape(item_url)
	m := Encrypt(url_query, 2, true)
	fmt.Println(m)
	mr := Reverse(m)
	ms := Split(mr, 2)
	str := Conv(ms)
	//	d := Decrypt(m,mr,false)
	fmt.Println(str)
	return str
}

//字符串反转
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
//字符串反转
func Split(s string, with int) []string {
	//	runes := []rune(s)
	var arr  []string
	for i := 0; i < len(s) - with; i = i + with {
		si := s[i:i + with]
		arr = append(arr, si)
	}
	return arr
}
//遍历msg的字符,转换成相应with位16进制,然后连接在一起
func Encrypt(msg string, with int, reverse bool) string {
	var ch_arr []string
	for _, ch := range msg {
		ch_int := int(ch)
		ch_arr = append(ch_arr, ToHex(ch_int + 88, with))
	}
	ret := strings.Join(ch_arr, "")
	if reverse {
		return Reverse(ret)
	}
	return ret
}
//遍历msg的字符,转换成相应with位16进制,然后连接在一起
func Decrypt(msg string, with int, reverse bool) string {
	var ch_arr []string
	for _, ch := range msg {
		ch_int := int(ch)
		ch_arr = append(ch_arr, ToHex(ch_int + 88, with))
	}
	ret := strings.Join(ch_arr, "")
	if reverse {
		return Reverse(ret)
	}
	return ret
}
//变成对应位数的16进制
func ToHex(ch int, width int) string {
	return fmt.Sprintf("%0" + strconv.Itoa(width) + "x", ch)
}//变成对应位数的16进制
func Conv(str_arr []string) string {
	var str string
	for _, v := range str_arr {
		n, _ := strconv.ParseUint(v, 16, 64)
		str += string(n - 88)
	}
	return str
}

func hexdec(s string) uint64 {
	d := uint64(0)
	for i := 0; i < len(s); i++ {
		x := uint64(s[i])
		if x >= 'a' {
			x -= 'a' - 'A'
		}
		d1 := x - '0'
		if d1 > 9 {
			d1 = 10 + d1 - ('A' - '0')
		}
		if 0 > d1 || d1 > 15 {
			panic("hexdec")
		}
		d = (16 * d) + d1
	}
	return d
}