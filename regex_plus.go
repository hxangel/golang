package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"time"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

type HtmlParse struct {
	url      string
	content  string
	replaces [][]string
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
var msg string

func NewHtmlParse() *HtmlParse {
	return &HtmlParse{
		replaces: [][]string{
			{`\s+`, " "}, //过滤多余回车
			{`<[ ]+`, "<"}, //过滤<__("<"号后面带空格)
			{`<\!–.*?–>`, ""}, // //注释
			{`<(\!.*?)>`, ""}, //过滤DOCTYPE
			{`<(\/?html.*?)>`, ""}, //过滤html标签
			{`<(\/?br.*?)>`, ""}, //过滤br标签
			{`<(\/?head.*?)>`, ""}, //过滤head标签
			{`<(\/?meta.*?)>`, ""}, //过滤meta标签
			{`<(\/?body.*?)>`, ""}, //过滤body标签
			{`<(\/?link.*?)>`, ""}, //过滤link标签
			{`<(\/?form.*?)>`, ""}, //过滤form标签
			{`<(applet.*?)>(.*?)<(\/applet.*?)>`, ""}, //过滤applet标签
			{`<(font.*?)>(.*?)<(\/font.*?)>`, ""}, //过滤applet标签
			{`<(\/?applet.*?)>`, ""},
			{`<(style.*?)>(.*?)<(\/style.*?)>`, ""}, //过滤style标签
			{`<(\/?style.*?)>`, ""},
			{`<(title.*?)>(.*?)<(\/title.*?)>`, ""}, //过滤title标签
			{`<(\/?title.*?)>`, ""},
			{`<(object.*?)>(.*?)<(\/object.*?)>`, ""}, //过滤object标签
			{`<(\/?objec.*?)>`, ""},
			{`<(noframes.*?)>(.*?)<(\/noframes.*?)>`, ""}, //过滤noframes标签
			{`<(\/?noframes.*?)>`, ""},
			{`<(i?frame.*?)>(.*?)<(\/i?frame.*?)>`, ""}, //过滤frame标签
			{`<(script.*?)>(.*?)<(\/script.*?)>`, ""}, //过滤script标签
			{`<(\/?script.*?)>`, ""},
			{`<(noscript.*?)>(.*?)<(\/noscript.*?)>`, ""}, //过滤noframes标签
			{`on([a-z]+)\s*="(.*?)"`, ""}, //过滤dom事件
			{`on([a-z]+)\s*='(.*?)'`, ""},
		},
	}
}

func (hp *HtmlParse) LoadUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println("HTTP状态码：", resp.StatusCode)
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

func mysqlInit() mysql.Conn {
	user := "root"
	pass := "root"
	db_name := "wmsys"
	//proto  := "unix"
	//addr   := "/var/run/mysqld/mysqld.sock"
	protocol := "tcp"
	addr := "localhost:3306"
	db := mysql.New(protocol, "", addr, user, pass, db_name)
	checkError(db.Connect())
	RightLine("数据库连接成功")
	return db
}

func checkError(err error) {
	fmt.Println(err);
	if err != nil {
		fmt.Println(err)
		//修改os.Exit(1)为return，因为会导致程序退出->退出返回值为1
		fmt.Println(time.Now().String() ,err)
		msg = "检查错误"
		ErrorLine(msg)
		return
	}
}
func ErrorLine(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行: " + msg)
	}
}
func RightLine(msg string) {
	fmt.Println(time.Now().String(),msg)
//	_, file, line, ok := runtime.Caller(1)
//	if ok {
//		fmt.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行" + "正常")
//	}
}
func main() {

	db := mysqlInit()
//	_, res, err := db.Query("select * from cps_mmb_order_info")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(res.Fields())
//	db.Close();

	/**************************************************************/
	hp := NewHtmlParse()
	err := hp.LoadUrl("http://qudao.ebinf.com/mmb-union/cpsStat.jsp?uuuu=jinliquan&pppp=maiM_17Gou&startDate=2014-07-10")
	if err != nil {
		fmt.Println(err)
	}
	// hp.Clear()
	//	fmt.Println(hp.content)
	m := hp.FindByTagName("tr")
	//	fmt.Println(len(m[586]))
	//	fmt.Println(' ')
	var q string = " INSERT INTO `cps_mmb_order_info` (`order_id`,`order_price`,`order_status`,`share_status`,`share_price`,`settlement_status`,`goods_type_id`,`channel_id`,`created_time`) VALUES "
	//	(1, 1, 5, 0),
	//	(2, 1, 9, 1),
	//	fmt.Println(m[0][2])
	var d string
	for i, _ := range m {
		mapx := getOrderInfo(m[i][2])
		if i == 0 {
			continue ;
		}
		var qstr string = fmt.Sprintf(`('%s','%s','%s','%s','%s','%s','%s','%s','%s')`, mapx["order_id"], mapx["order_price"], mapx["order_status"], mapx["share_status"], mapx["share_price"], mapx["settlement_status"], mapx["goods_type_id"], mapx["channel_id"], mapx["created_time"]);
		if d !=""{
			d =d+","+qstr
		}else{
			d = qstr
		}

		//		fmt.Println(i, mapx["order_id"]);

	}
	q = q+d+";"
	_,_,dberr := db.Query(q);
	fmt.Println(dberr);
	fmt.Println(db.Close())
	fmt.Println(q);

	//	fmt.Println(' ')
	//	fmt.Println(m)
	/**************************************************************/

	/**************************************************************/
	/*hp := NewHtmlParse()
	html := "<div><tr class='a'><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr></div>"
	hp.LoadStr(html)
	m := hp.FindByAttr("tr", "class", "a")
	fmt.Println(len(m[:1][0][2]))
	fmt.Println(m[:1][0][2])*/
	/**************************************************************/

	/**************************************************************/
	/*hp := NewHtmlParse()
	html := "<tr id='a' class=aaa><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>"
	hp.LoadStr(html)
	m := hp.FindByAttr("tr", "id", "a")
	fmt.Println(m)*/
	/**************************************************************/

	/**************************************************************/
	/*hp := NewHtmlParse()
	html := `<tr><td class="a" a="rr">aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>`
	hp.LoadStr(html)
	m := hp.FindByAttr("td", "class", "a")
	fmt.Println(m)*/
	/**************************************************************/
}
