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
var msg string
var num int

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
	//	fmt.Println(re.String())
	//	return re.FindAllStringSubmatch(hp.content, -1)
	return re.FindAllStringSubmatch(hp.content, -1)
}

func (hp *HtmlParse) GetAttrValue(attr, rule string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)%s=['"]%s['"]).*?`, attr, rule,))
	//	fmt.Println(re.String())
	//fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
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

func SaveFile(url, pathstr string) error {
	num +=1
	file := "/www/htdocs/metronic"
	abs := fmt.Sprintf(`%s/%s`, file, pathstr)
	fmt.Println(abs)
	_, err := exec.LookPath(abs)
	//
	//	ext := path.Ext(pathstr)
	//	mp := make(map[string]string)
	//	mp[".jpg"] = "jpg"
	//	mp[".png"] = "png"
	//	mp[".jif"] = "jif"
	//	if _, ok := mp[ext]; ok {
	//		fmt.Println(abs)
	//	}
	if err != nil {
		fmt.Println(abs)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			fmt.Println(url)
			fmt.Println("HTTP状态码：", resp.StatusCode)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		//		msg := fmt.Sprintf("%s is unavailable \n", abs)
		//		fmt.Print(len(msg))
		pathstr = path.Dir(abs)
		//		fmt.Printf("creating %s \n", f)
		os.MkdirAll(pathstr, 0777)
		//		fmt.Printf("writing %s \n", f)
		ioutil.WriteFile(abs, body, 0777)
		return nil

	}
	//测试环境为ArchLinux
	//	fmt.Printf("%s is available \n", abs)

	return nil
}

func DealStatic(files [][]string) {
	for i, _ := range files {
		url := baseurl + files[i][2]
		SaveFile(url, "metronic/"+files[i][2])
	}
}

func Anycss(files [][]string) {
	file, _ := os.Getwd()
	re := regexp.MustCompile(`url\("(.*)"\)`)
	for i, _ := range files {
		abs := fmt.Sprintf(`%s/metronic/%s`, file, files[i][2])
		_, err := exec.LookPath(abs)
		url := baseurl + files[i][2]
		if err != nil {
			resp, err := http.Get(url)
			if err != nil {
				println("%s", err)
			}
			if resp.StatusCode != 200 {

				fmt.Println("HTTP状态码：", resp.StatusCode)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			img := re.FindAllStringSubmatch(fmt.Sprintf("%s", body), -1)
			if img != nil {
				getBackGround(files[i][2], img)
			}
		}else {
			body, err := ioutil.ReadFile(abs)
			if err != nil {

			}
			img := re.FindAllStringSubmatch(fmt.Sprintf("%s", body), -1)
			if img != nil {
				getBackGround(files[i][2], img)
			}
		}

		//		SaveFile(url, "metronic/"+files[i][2])
	}
}

func getBackGround(css string, img [][]string) {
	p := path.Dir(css)
	for i, _ := range img {
		uri := fmt.Sprintf("%s/%s", p, img[i][1])
		url := baseurl + uri
		SaveFile(url, "metronic/"+uri)
	}
}
func main() {
	file, _ := os.Getwd()
	abs := fmt.Sprintf(`%s/xx.html`, file)
	re := regexp.MustCompile(`milo.com/(.*)\s404`)
	body, err := ioutil.ReadFile(abs)
	if err != nil {

	}
	files := re.FindAllStringSubmatch(fmt.Sprintf("%s", body), -1)
	for i, _:= range files {
		SaveFile(baseurl+files[i][1], files[i][1])
	}

}

func grap(filestr string, hp *HtmlParse) {
	url := baseurl + "templates/admin/" + filestr
	pathstr := "metronic/templates/admin/" + filestr
	SaveFile(url, pathstr)
	js := hp.GetAttrValue("src", "../../(.*.js)")
	DealStatic(js)
	css := hp.GetAttrValue("href", "../../(.*.css)")
	DealStatic(css)
	Anycss(css)
	img := hp.GetAttrValue("src", "../../(.*.png)")
	DealStatic(img)

}
func getLinks(url string) [][]string {
	hp := NewHtmlParse()
	err := hp.LoadUrl(url)
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println(hp.content)
	m := hp.GetAttrValue("href", "(.*.html)")

	for i, _ := range m {
		grap(m[i][2], hp)

	}
	return m
}

func mainbackup() {

	var inx = "http://www.keenthemes.com/preview/metronic/templates/admin/index.html"
	getLinks(inx);
	println(num)
	//	u, _ := url.Parse(link)
	//	fmt.Println(u)
	//	/**************************************************************/
	//	hp := NewHtmlParse()
	//	//	err := hp.LoadUrl("http://qudao.ebinf.com/mmb-union/cpsStat.jsp?uuuu=jinliquan&pppp=maiM_17Gou&startDate=2014-07-10")
	//	err := hp.LoadUrl(link)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	// hp.Clear()
	//		fmt.Println(hp.content)
	//	m := hp.FindByTagName("tr")
	//	fmt.Println(len(m[586]))

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