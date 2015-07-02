package grap
import (
	"strconv"
	"time"
	"fmt"
	"runtime"
	"log"
)

func ErrorLine(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行: " + msg)
	}
}
func RightLine(msg string) {
	log.Println(time.Now().String(), msg)
}

func checkError(err error) {
	fmt.Println(err);
	if err != nil {
		msg = "检查错误"
		ErrorLine(msg)
		return
	}
}