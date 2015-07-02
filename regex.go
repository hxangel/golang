package main

import (
	"fmt"
	"regexp"
)

type HtmlParse struct {
}

type HtmlNode struct {
	tagName string
	id      string
	class   string
	text    string
}

func main() {
	html := "<tr><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>"
	m := GetsByTagName(html, "tr")
	fmt.Println(m)

	/*html := "<tr id='a' class=aaa><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>"
	m := GetsById(html, "tr", "a")
	fmt.Println(m)*/

	/*html := `<tr><td class="a" a="rr">aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>`
	m := GetsByClass(html, "td", "a")
	fmt.Println(m)*/
}

func GetsByTagName(html, tagName string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s>(.*)</%s>).*?`, tagName, tagName))
	// fmt.Println(re.String())
	return re.FindAllStringSubmatch(html, -1)
}

func GetsById(html, tagName, id string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*id=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, id, tagName))
	//fmt.Println(re.String())
	return re.FindAllStringSubmatch(html, -1)
}

func GetsByClass(html, tagName, class string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*class=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, class, tagName))
	// fmt.Println(re.String())
	return re.FindAllStringSubmatch(html, -1)
}
