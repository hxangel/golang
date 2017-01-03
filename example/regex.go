//Go 提供内置的正则表达式。这里是 Go 中基本的正则相关功能的例子。
package main
import "bytes"
import "fmt"
import "regexp"
func main() {


	regex, _ := regexp.Compile(`(?Us)pageConfig = ({.*});`)
	id := regex.FindStringSubmatch(` var pageConfig = {
            compatible: true,
            product: {
                skuid: 1500970,
                name: '\u7f8e\u7684\uff08\u004d\u0069\u0064\u0065\u0061\uff09\u5927\u0031\u5339\u0020\u667a\u80fd\u4e91\u9664\u7532\u919b\u51b7\u6696\u53d8\u9891\u6302\u673a\uff08\u4eac\u4e1c\u68a6\u60f3\u5347\u7ea7\u7248\uff09\u0020\u624b\u673a\u0041\u0050\u0050\u63a7\u5236\u0020\u004b\u0046\u0052\u002d\u0032\u0036\u0047\u0057\u002f\u0057\u004a\u0043\u0041\u0033\u0040',
                skuidkey:'80B90C7CB0B9D7ED857635D110B48B80',
                href: '//item.jd.com/1500970.html',
                src: 'jfs/t1939/321/392123766/74291/73a0d9d5/5603db1dN309ca724.jpg',
                cat: [737,794,870],
                brand: 12380,
                pType: 1,
                isClosePCShow: false,
                venderId:1000001452,
                shopId:'1000001452',
                                commentVersion:'3461',                 specialAttrs:["isFlashPurchase-0","isCanVAT","isPickingGoods-0","isHaveYB","packType","isFitService","isCanUseDQ-1","isSelfService-0","isCanUseJQ-1","isWeChatStock-0","HYKHSP-0","isNSNGgoods-3","isOverseaPurchase-0","is7ToReturn-1","isOldChangeNewForJD","YYSLLZC-0"],
                recommend : [0,1,2,3,4,5,6,7,8,9],
                easyBuyUrl:"http://easybuy.jd.com/skuDetail/newSubmitEasybuyOrder.action",
                                colorSize: [{"Color":"大1匹","SkuId":1133068,"Size":"纯呼吸 变频【独立送风除甲醛】"},{"Color":"1.5匹","SkuId":1133441,"Size":"纯呼吸 定速【独立送风除甲醛】"},{"Color":"大1匹","SkuId":1298903,"Size":"京东英雄 单冷【强劲制冷】"},{"Color":"小1匹","SkuId":1298893,"Size":"京东英雄 定速【迅猛制冷制热】"},{"Color":"小1匹","SkuId":1298894,"Size":"京东英雄 单冷【强劲制冷】"},{"Color":"小1.5匹","SkuId":1298898,"Size":"京东英雄 定速【迅猛制冷制热】"},{"Color":"1.5匹","SkuId":1298907,"Size":"京东英雄 变频【稀土压缩机】"},{"Color":"1.5匹","SkuId":1338105,"Size":"京东英雄 单冷【强劲制冷】"},{"Color":"小1.5匹","SkuId":1462425,"Size":"京东英雄 变频【稀土压缩机】"},{"Color":"大1匹","SkuId":1500970,"Size":"智能WIFI变频【独立送风除甲醛】"},{"Color":"1.5匹","SkuId":1500981,"Size":"智能WIFI变频【独立送风除甲醛】"}],                warestatus: 1,                                 tips: [{"order":3,"tip":"支持7天无理由退货"}],                isOtc: false,                isBookMvd4Baby: false,                                desc: '//d.3.cn/desc/1500970?cdn=2',                foot: '//d.3.cn/footer?type=common_config2'
            }
        };
                        try {
                        function is_sort_black_list() {
              var jump_sort_list = {"6881":3,"1195":3,"10011":3,"6980":3,"12360":3};
              if(jump_sort_list['737'] == 1 || jump_sort_list['794']==2 || jump_sort_list['870']==3) {
                return true;
              }
              return false;
            }`)
	fmt.Println(id[1])
	return
	//		http://m.jd.com/product/1143562.html
	//这个测试一个字符串是否符合一个表达式。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	fmt.Println(r.MatchString("peach"))
	//这是查找匹配字符串的。
	fmt.Println(r.FindString("peach punch"))
	//这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容。
	fmt.Println(r.FindStringIndex("peach punch"))
	//Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
	fmt.Println("FindStringSubmatch",r.FindStringSubmatch("peach punch"))
	//类似的，这个会返回完全匹配和局部匹配的索引位置。
	fmt.Println("FindStringSubmatchIndex",r.FindStringSubmatchIndex("peach punch"))
	//带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	//All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	//这个函数提供一个正整数来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	//上面的例子中，我们使用了字符串作为参数，并使用了如 MatchString 这样的方法。我们也可以提供 []byte参数并将 String 从函数命中去掉。
	fmt.Println(r.Match([]byte("peach")))
	//创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	//regexp 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	//Func 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
//$ go run regular-expressions.go
//true
//true
//peach
//[0 5]
//[peach ea]
//[0 5 1 3]
//[peach punch pinch]
//[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
//[peach punch]
//true
//p([a-z]+)ch
//a <fruit>
//a PEACH
//完整的 Go 正则表达式参考，请查阅 regexp 包文档。
