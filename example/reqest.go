package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	surl := fmt.Sprintf("https://s.taobao.com/search?q=%s", "2015冬新款真皮雪地靴平底休闲短靴女潮流苏保暖短筒靴女棉鞋靴子")

	resp, err := http.Get(surl)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}