package main

import (
	"github.com/parnurzeal/gorequest"
	"fmt"
)

func main() {
	//pollURL := "http://zhushou.huihui.cn/productSense?av=3.0&browser=chrome&extensionid=eb5a2408-3e2b-6d03-1007-f40d2c9120df&k=00cc00956dcf97aa8535008a00880089008f667d891d66086b967f0782ca4e5b525e88ee4e6288bb4e7c4f4e59af6c6c8d808124601b8931664e5c7259af891d008f0090008a008a00b6007e00c3009500c400d000cb00d000b6007e00bc009500c400cb&m=19d8888809b8d8a80988c8d8c9b8d7cb1ce9b8d75ccc0c685cdbcc1ce9a8d75c7cbb684c4c9b5ccc684c1c9bccdbcbe9a8d7e9a8d799b8d7bc8ccccc0c&t=1491634931676552823&vendor=chromenew&version=4.2.9.1"
	pollURL := "http://golang.org"
	_, bodyString, errs := gorequest.New().Post(pollURL).Send().End()
	if errs != nil {
		fmt.Println(errs)
		return
	}
	fmt.Println(bodyString)

}
