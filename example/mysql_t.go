//数据库连接池测试
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gou?charset=utf8")
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(15)
	db.Ping()
}

func main() {
	start := time.Now()
	rows, err := db.Query("select * FROM gou_third_goods_price WHERE  `goods_id`  = '547466'")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	time.Sleep(time.Second*5)
	fmt.Println(record)
	end := time.Now()
	fmt.Println(end)
	//输出执行时间，单位为毫秒。
	fmt.Println(end.Sub(start).Nanoseconds() / 1000000)
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func pool(w http.ResponseWriter, r *http.Request) {

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
