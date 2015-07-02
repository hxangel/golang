package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"time"
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

type StaticFile struct{
	url       string
	path      string
	realpath  string
	body      []byte
}

type HtmlNode struct {
	tagName string
	id      string
	class   string
	text    string
}

// embed regexp.Regexp in a new type so we can extend it
type myRegexp struct {
	*regexp.Regexp
}

var baseurl = "http://www.keenthemes.com/preview/metronic/"
var basepath = "/www/htdocs/met/"
var msg string
var num int

func NewHtmlParse() *HtmlParse {
	return &HtmlParse{
		replaces: [][]string{
		},
	}
}
func StaticFileParse() *StaticFile {
	return &StaticFile{
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
func (sp *StaticFile) GetAttrValue(attr, rule string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)%s\s?=\s?['|"]%s).*?`, attr, rule,))
	//	fmt.Println(re.String())
	return re.FindAllStringSubmatch(fmt.Sprintf("%s", sp.body), -1)
}
func (hp *HtmlParse) FindByAttr(tagName, attr, value string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*%s=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, attr, value, tagName))
	//fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}

func getOrderInfo(str string) map[string]string {
	//	fmt.println(str);
	myre := myRegexp{regexp.MustCompile(fmt.Sprintf(`((?U))<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>\s<td>(?P<%s>.*)</td>`, "serial_id", "order_id", "order_price", "realname", "phone", "address", "order_status", "share_status", "share_price", "settlement_status", "goods_type_id", "channel_id", "created_time"))}
	return myre.FindStringSubmatchMap(str);

}

func (r *myRegexp) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}
	for i, name := range r.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}

		captures[name] = match[i]

	}
	return captures
}

func checkError(err error) {
	fmt.Println(err);
	if err != nil {
		fmt.Println(err)
		//修改os.Exit(1)为return，因为会导致程序退出->退出返回值为1
		fmt.Println(time.Now().String(), err)
		msg = "检查错误"
		ErrorLine(msg)
		return
	}
}


func getFileSaveName(u string) {
	fmt.Println(u);
}

func (sp *StaticFile) ReadFileFromUri(RelativePath string) error {
	AbsolutePath := fmt.Sprintf(`%s%s`, basepath, RelativePath)
	url := fmt.Sprintf(`%s%s`, baseurl, RelativePath)
	_, err := exec.LookPath(AbsolutePath)
	if err != nil {
		resp, err := http.Get(url)
		if err != nil {
			println("%s", err)
		}
		if resp.StatusCode != 200 {
			println("HTTP状态码：%s", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		sp.body = body
		if err != nil {
			println("																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																	%s", err)
		}
		if err != nil {
			return err
		}
		//		fmt.Println(url)
		//		println(AbsolutePath)
		//		msg := fmt.Sprintf("%s is unavailable \n", AbsolutePath)
		//		fmt.Print(len(msg))
		FileDir := path.Dir(AbsolutePath)
		os.MkdirAll(FileDir, 0777)
		ioutil.WriteFile(AbsolutePath, body, 0777)
	}else {
		body, err := ioutil.ReadFile(AbsolutePath)
		if err != nil {
			println("%s", err)
		}
		sp.body = body
	}
	sp.path = RelativePath
	sp.realpath = AbsolutePath
	sp.url = url
	return nil

}
func (sp *StaticFile) AnyCss(files [][]string) {
	re := regexp.MustCompile(`url\("(.*)"\)`)
	for i, _ := range files {
		if path.Ext(files[i][2]) == "css" {
			err := sp.ReadFileFromUri(files[i][2])
			checkError(err)
			img := re.FindAllStringSubmatch(fmt.Sprintf("%s", sp.body), -1)
			if img != nil {
				sp.getBackGround(files[i][2], img)
			}
		}
	}
}

func (sp *StaticFile) getBackGround(css string, img [][]string) {
	p := path.Dir(css)
	for i, _ := range img {
		uri := fmt.Sprintf("%s/%s", p, img[i][1])
		sp.ReadFileFromUri(uri)
	}
}

func Claw(HtmlFileName string) {
	//	fmt.Println(HtmlFileName)
	sp := StaticFileParse()
	RelativePath := "templates/admin/" + HtmlFileName
	_ = sp.ReadFileFromUri(RelativePath)

	//	fmt.Println(sf.url)
	static := sp.GetAttrValue(`[src|href]+`, `../../(.*.jpg|.*.jpeg|.*.gif|.*[.]+css|.*.png|.*[.]+js|.*.ico|.*.swf|.*.eot|.*.svg|.*.ttf|.*.woff)`)
	for i, _ := range static {
		_ = sp.ReadFileFromUri(static[i][2])
		//		fmt.Println(sp.url)
	}

	sp.AnyCss(static)
}

func main() {

	var inx = "http://www.keenthemes.com/preview/metronic/templates/admin/index.html"
	hp := NewHtmlParse()
	err := hp.LoadUrl(inx)
	if err != nil {
		fmt.Println(err)
	}
	m := hp.GetAttrValue("href", `([^http].*html)`)
	//	Claw(m[15][2])
	for i, _ := range m {
		Claw(m[i][2])

	}

}


func ErrorLine(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行: " + msg)
	}
}
func RightLine(msg string) {
	fmt.Println(time.Now().String(), msg)
}
