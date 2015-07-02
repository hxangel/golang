package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	//"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var c = make(chan int)
var msg string
func scanDir(w http.ResponseWriter, d string) {
	dir, err := os.Open(d)
	if err != nil {
		// panic(nil)太粗暴，容易让程序崩溃，改成return，并打印错误行数
		msg = "打开目录失败"
		ErrorLine(msg)
		return
	}
	defer dir.Close()
	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		// panic(err)太粗暴，容易让程序崩溃，改成return，并打印错误行数
		msg = "读取目录失败"
		ErrorLine(msg)
		return
	}

	db := mysqlInit()
	// 遍历文件列表
	runtime.GOMAXPROCS(4)
	i := 0
	for _, fi := range fis {
		// 逃过文件夹, 我这里就不递归了
		fi.Name()
		fmt.Println(time.Now().String() + fi.Name())

		if fi.IsDir() {
			RightLine()
			continue
		}
		// 打印文件名称
		full_path := d + "/" + fi.Name()
		i++

		//fmt.Ffmt.Printlnf(w, "正在扫描%s \n", full_path )
		RightLine()
		go scanApk(w, db, full_path)
	}
	//fmt.Printf("一共%d \n", i )
	for m := 0; m < i; m++ {
		RightLine()

		<-c
		//fmt.Printf("收到%d \n", m)
	}
	//增加关闭数据库
	db.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		//修改os.Exit(1)为return，因为会导致程序退出->退出返回值为1
		fmt.Println(time.Now().String() + "checkError函数检查出错误")
		msg = "检查错误"
		ErrorLine(msg)
		return
	}
}

func queryByPath(w http.ResponseWriter, db mysql.Conn, pt string) {
	rows, res := checkedResult(db.Query("select full_result from risks_check_history where apk='%s' ORDER BY id DESC Limit 1", pt))
	full_result := res.Map("full_result")
	for _, row := range rows {
		//fmt.Ffmt.Printlnf(w,
		fmt.Fprintf(w,
			"------[ %s 包含以下风险]---------\n%s\n",
			pt,
			row.Str(full_result),
		)
		RightLine()
	}
}

func scanApk(w http.ResponseWriter, db mysql.Conn, pt string) {
	run_risk := "/var/www/deployment/run.sh " + pt
	err := exec.Command("sh", "-c", run_risk).Run()
	if err != nil {
		msg = "执行run.sh错误"
		ErrorLine(msg)
		//修改log.Fatal(err)为return，因为里log模块Fatal函数会导致程序调用os.Exit(1)退出->退出返回值为1
		fmt.Println(time.Now().String() + "scanApk函数检查出致命错误")
		return
	}
	queryByPath(w, db, pt)
	RightLine()
	c <- 1
}

func checkedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row, mysql.Result) {
	RightLine()
	checkError(err)
	return rows, res

}

func mysqlInit() mysql.Conn {
	user := "apk"
	pass := "apk"
	dbname := "apk_risk"
	//proto  := "unix"
	//addr   := "/var/run/mysqld/mysqld.sock"
	proto := "tcp"
	addr := "127.0.0.1:3306"
	db := mysql.New(proto, "", addr, user, pass, dbname)
	RightLine()
	checkError(db.Connect())
	RightLine()
	return db
}

func handlerScan(w http.ResponseWriter, r *http.Request) {

	orgi_dir := "/data/ad/scan"
	RightLine()
	scanDir(w, orgi_dir)

}

func ErrorLine(msg string ) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行: "+msg )
	}
}

func RightLine() {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(time.Now().String() + file + "第" + strconv.Itoa(line) + "行"+"正常")
	}
}

func main() {
	http.HandleFunc("/scan", handlerScan)
	RightLine()
	http.ListenAndServe(":999", nil)

}
