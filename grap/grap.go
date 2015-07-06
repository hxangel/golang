package grap

import (
	"log"
	"fmt"
)

type HtmlNode struct {
	tagName string
	id      string
	class   string
	text    string
}
//抓取html页面
func Claw(HtmlFileName string) {

	//	fmt.Println(HtmlFileName)
	sp := StaticFileParse()
	RelativePath := "templates/admin/" + HtmlFileName
	//	fmt.Println(RelativePath)
	_ = sp.ReadFileFromUri(RelativePath)

	//	fmt.Println(sf.url)
	static := sp.GetAttrValue(`[src|href]+`, `../../(.*.jpg|.*.jpeg|.*.json|.*.gif|.*[.]+css|.*.png|.*[.]+js|.*.ico|.*.swf|.*.eot|.*.svg|.*.ttf|.*.woff|.*.woff2)`)
	for i, _ := range static {
	//	fmt.Println(static[i][2])
		_ = sp.ReadFileFromUri(static[i][2])
		//		fmt.Println(sp.url)
	}

	sp.AnyCss(static)
}

func Grap(idx_url string ) {
	fmt.Println("Start")
	hp := NewHtmlParse()
	err := hp.LoadUrl(idx_url)
	if err != nil {
		log.Println(err)
	}
	m := hp.GetAttrValue("href", `(.*html)`)
	//	Claw(m[15][2])
	for i, _ := range m {
		fmt.Println(m[i][2])
		Claw(m[i][2])
	}
	fmt.Println("Finish")
}