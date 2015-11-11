package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

type QueryResp struct {
	Addr string
	Time float64
}

func query(ip string, port string, c chan QueryResp) {
	start_ts := time.Now()
	var timeout = time.Duration(30 * time.Second)
	host := fmt.Sprintf("%s:%s", ip, port)
	url_proxy := &url.URL{Host: host}
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(url_proxy)},
		Timeout:   timeout}
	resp, err := client.Get("http://err.taobao.com/error1.html")
	if err != nil {
		c <- QueryResp{Addr: host, Time: float64(-1)}
		return
	}

	if resp.StatusCode !=200 {
		c <- QueryResp{Addr: host, Time: float64(-1)}
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	time_diff := time.Now().UnixNano() - start_ts.UnixNano()
	if strings.Contains(string(body), "alibaba.com") {
		c <- QueryResp{Addr: host, Time: float64(time_diff) / 1e9}
	} else {
		c <- QueryResp{Addr: host, Time: float64(-1)}
	}
}

func main() {

	dats := map[string]string{
		"36.61.75.116":"3128",
		"112.195.84.134":"9000",
		"223.151.37.134":"9000",
		"120.195.201.73":"80",
		"218.92.227.172":"34043",
		"120.195.205.203":"80",
		"120.195.202.243":"80",
		"120.195.199.248":"80",
		"220.181.32.106":"80",
		"122.226.142.52":"3128",
	}
	runtime.GOMAXPROCS(4)

	resp_chan := make(chan QueryResp, 10)

	for ip, port := range dats {
		go query(ip, port, resp_chan)
	}

	for i, j := range dats {
		fmt.Println(i, j)
		r := <-resp_chan
		if r.Time > 0 {
			fmt.Printf("%s %v\n", r.Addr, r.Time)
		}
	}
}