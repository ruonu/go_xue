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

func main() {
	db, err := sql.Open("mysql", "test:123456@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	fmt.Println("开始执行数据查找")

	//搜索数据
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var time1 string
		err = rows.Scan(&uid, &username, &department, &time1)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(time1)
	}

	//输出测试结果
	fmt.Println("测试一下json输出")

	//Getdict_mysql()

	b, err := json.Marshal(rows)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)

	//执行查询操作

	rows2, err := db.Query("SELECT * FROM `user` WHERE `id` > 0")
	if err != nil {

		fmt.Println("select db failed,err:", err)

		return

	}
	// 这里获取的rows是从数据库查的满足user_id>=5的所有行的email信息，rows.Next(),用于循环获取所有
	fmt.Println("新的查询方式")

	for rows2.Next() {

		var s string

		err = rows2.Scan(&s)
		if err != nil {

			fmt.Println(err)

			return

		}

		fmt.Println(rows2)

	}

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
