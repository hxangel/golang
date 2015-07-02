package grap

import (
	"fmt"
	"log"
)

type HtmlNode struct {
	tagName string
	id      string
	class   string
	text    string
}

func Claw(HtmlFileName string) {
	//	fmt.Println(HtmlFileName)
	sp := StaticFileParse()
	RelativePath := "templates/admin/" + HtmlFileName
	fmt.Println(RelativePath)
	_ = sp.ReadFileFromUri(RelativePath)

	//	fmt.Println(sf.url)
	static := sp.GetAttrValue(`[src|href]+`, `../../(.*.jpg|.*.jpeg|.*.gif|.*[.]+css|.*.png|.*[.]+js|.*.ico|.*.swf|.*.eot|.*.svg|.*.ttf|.*.woff)`)
	for i, _ := range static {
		fmt.Println(static[i][2])
		_ = sp.ReadFileFromUri(static[i][2])
		//		fmt.Println(sp.url)
	}

	sp.AnyCss(static)
}

func Grap(idx_url string ) {

	hp := NewHtmlParse()
	err := hp.LoadUrl(idx_url)
	if err != nil {
		log.Println(err)
	}
	m := hp.GetAttrValue("href", `([^http].*html)`)
	//	Claw(m[15][2])
	for i, _ := range m {
		fmt.Println(m[i][2])
		Claw(m[i][2])
	}

}