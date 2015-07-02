package grap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HtmlParse struct {
	url       string
	uri       string
	script    [][]string
	style     [][]string
	img       [][]string
	content   string
	replaces  [][]string
}



var msg string

func NewHtmlParse() *HtmlParse {
	return &HtmlParse{
		replaces: [][]string{
		},
	}
}


func (hp *HtmlParse) LoadUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//	fmt.Println("HTTP状态码：", resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	hp.content = fmt.Sprintf("%s", body)
	hp.Clear()
	return nil
}

func (hp *HtmlParse) LoadStr(content string) {
	hp.content = content
}

func (hp *HtmlParse) Clear() {
	length := len(hp.replaces)
	for i := 0; i < length; i++ {
		if l := len(hp.replaces[i]); l > 0 {
			p, r := hp.replaces[i][:1], hp.replaces[i][1:2]
			hp.content = regexp.MustCompile(p[0]).ReplaceAllString(hp.content, r[0])
		}
	}
}

func (hp *HtmlParse) FindByTagName(tagName string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*>(.*)</%s>).*?`, tagName, tagName))
	return re.FindAllStringSubmatch(hp.content, -1)
}

func (hp *HtmlParse) GetAttrValue(attr, rule string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)%s\s?=\s?['|"]%s).*?`, attr, rule,))
	//	fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}

func (hp *HtmlParse) FindByAttr(tagName, attr, value string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*%s=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, attr, value, tagName))
	//fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}






func getFileSaveName(u string) {
	fmt.Println(u);
}


