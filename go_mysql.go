package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type user struct {
	Uid        int    `json:"uid"`
	UserName   string `json:"user_name"`
	Department string `json:"department"`
	Time1      string `json:"time_1"`
}

func main() {
	db, err := sql.Open("mysql", "test:123456@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	fmt.Println("开始执行数据查找")

	//搜索数据
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	var data []user
	for rows.Next() {

		info := user{}
		err = rows.Scan(&info.Uid, &info.UserName, &info.Department, &info.Time1)
		checkErr(err)
		data = append(data, info)
		fmt.Println(info)
	}

	//输出测试结果
	fmt.Println("测试一下json输出")

	//Getdict_mysql()

	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Getdict_mysql() {

	return

}
