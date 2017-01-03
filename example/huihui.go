package main
import (
	"fmt"
	"strings"
	"net/url"
	"strconv"
	"time"
	"github.com/hxangel/spider"
	"encoding/json"
)


func main() {
	item_url := "http://item.jd.com/1510479.html"
	title := "创维(Skyworth) 42E5ERS 42英寸 高清LED窄边平板液晶电视(银色)"
	url_str := huihui(item_url, title)
	fmt.Println(url_str)

	loader := spider.NewLoader()
	content, err := loader.Send(url_str, "Get", nil)
	if err != nil {
		return
	}

	//	解析json
	var data_json map[string]interface{}
	if err := json.Unmarshal(content, &data_json); err != nil {
		return
	}
	//	判断状态
	thisItem := data_json["thisItem"].(map[string]interface{})
	list     := data_json["urlPriceList"].([]interface{})

	for _,val := range list{
		row :=val.(map[string]interface{})
		item :=row["items"].([]interface{})[0].(map[string]interface{})
		inf := spider.Sense{Title:item["name"].(string),Price:item["price"].(string),ItemUrl:item["url"].(string)}
		inf.GetChannelBySite(row["site"].(string))
		inf.GetItemID(item["url"].(string))
		inf.GetHistoryPrice()
		fmt.Println(inf)
	}

	fmt.Println(thisItem["price"].(float64))
}










func huihui(item_url string, title string) string {
	url_query := url.QueryEscape(item_url)
	m := encrypt(url_query, 2, true)
	title_map := []string{"t=" + title, "k=lxsx", "d=ls"}
	title_param_str := strings.Join(title_map, "^&")
	k := encrypt(title_param_str, 4, false)

	parameters := url.Values{}
	parameters.Add("m", m)
	parameters.Add("k", k)
	parameters.Add("av", "3.0", )
	parameters.Add("vendor", "chrome", )
	parameters.Add("browser", "chrome")
	parameters.Add("version", "3.7.5.2", )
	parameters.Add("extensionid", "e8170cff-a3e7-039c-3865-44b1c126227e", )
	parameters.Add("t", fmt.Sprintf("%d", time.Now().UnixNano()))

	var Url *url.URL
	Url, _ = url.Parse("http://zhushou.huihui.cn/productSense")
	Url.RawQuery = parameters.Encode()
	return Url.String()
}



func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func encrypt(msg string, with int, reverse bool) string {
	var s_arr []string
	for _, s_chr := range msg {
		s_int := int(s_chr)
		s_arr = append(s_arr, to_hex(s_int, with))
	}
	ret := strings.Join(s_arr, "")
	if reverse {
		return Reverse(ret)
	}
	return ret
}


func to_hex(c int, width int) string {
	return fmt.Sprintf("%0" + strconv.Itoa(width) + "x", c + 88)
}
