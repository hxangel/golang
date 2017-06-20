package main

import (
	"github.com/hxangel/spider"
)

func main() {
	hh := &spider.Hhui{}
	//item_url := "http://item.jd.com/1510479.html"
	title := "创维(Skyworth) 42E5ERS 42英寸 高清LED窄边平板液晶电视(银色)"
	params := map[string]string{"id": "1510479", "channel":"jd","title":title, "callback": ""}
	var i = &spider.Item{Params:params}
	hh.Item(i)
}
