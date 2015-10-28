package main
import (
	"fmt"
	"net/http"
)
func main() {
	url := "http://www.suning.com/emall/sprd_10052_10051_21318695_0070074646_.html"
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer resp.Body.Close()
	fmt.Println("Header:", resp.Request.URL)
}