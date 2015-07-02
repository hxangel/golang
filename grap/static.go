package grap

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"io/ioutil"
	"net/http"
	"regexp"
)

type StaticFile struct{
	url       string
	path      string
	realpath  string
	body      []byte
}

// embed regexp.Regexp in a new type so we can extend it
type myRegexp struct {
	*regexp.Regexp
}

var baseurl = "http://www.keenthemes.com/preview/metronic/theme/"
var basepath = "/home/milo/Documents/met/theme/"


func StaticFileParse() *StaticFile {
	return &StaticFile{
	}
}





func (sp *StaticFile) GetAttrValue(attr, rule string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)%s\s?=\s?['|"]%s).*?`, attr, rule,))
	return re.FindAllStringSubmatch(fmt.Sprintf("%s", sp.body), -1)
}

func (sp *StaticFile) ReadFileFromUri(RelativePath string) error {
	AbsolutePath := fmt.Sprintf(`%s%s`, basepath, RelativePath)
	url := fmt.Sprintf(`%s%s`, baseurl, RelativePath)
	fmt.Println(url)
	_, err := exec.LookPath(AbsolutePath)
	if err != nil {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Sprintf("%s", err)
		}
		if resp.StatusCode != 200 {
			fmt.Sprintf("HTTP状态码：%s", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		sp.body = body
		if err != nil {
			println("																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																	%s", err)
		}
		if err != nil {
			return err
		}
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

