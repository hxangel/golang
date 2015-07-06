package grap
import (
	"strconv"
	"time"
	"fmt"
	"runtime"
	"log"
	"path"
	"os"
)
var error_log  ="/www/dev/go/error.log"
var access_log ="/www/dev/go/access.log"
func ErrorLine(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行: " + msg)
	}
}
func RightLine(msg string) {
	log.Println(time.Now().String(), msg)
}

func Error(content string) {
	FileDir := path.Dir(error_log)
	os.MkdirAll(FileDir, 0777)
	write(error_log, content)
}

func Right(content string) {
	FileDir := path.Dir(access_log)
	os.MkdirAll(FileDir, 0777)
	write(access_log, content)
}

func write(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	text := fmt.Sprintln(content)
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func checkError(err error) {
	fmt.Println(err);
	if err != nil {
		msg = "检查错误"
		ErrorLine(msg)
		return
	}
}