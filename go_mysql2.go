package main

/**
Author:ShaZhenYu
Date:20180926
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_user   = "test"
	mysql_passwd = "123456"
	mysql_ip     = "127.0.0.1"
	mysql_port   = "3306"
	mysql_robot  = "test"
	mysql_qyDB   = "test"
)

func main() {
	//此示例传入一个参数，返回两个参数
	var qID = "0"
	//读取mysql指定表的指定字段内容
	tmpAnswer, tmpCodeName := funReadSql(qID)
	//读取mysql成功
	fmt.Printf("读取mysql成功！！！\ncodeName:%s, answer:%s\n", tmpCodeName, tmpAnswer)
}

func funReadSql(quitionID string) (string, string) {
	//打开数据库
	db, errOpen := sql.Open("mysql", mysql_user+":"+mysql_passwd+"@tcp("+mysql_ip+":"+mysql_port+")/"+mysql_qyDB+"?charset=utf8")
	if errOpen != nil {
		//TODO，这里只是打印了一下，并没有做异常处理
		fmt.Println("funReadSql Open is error")
	}

	//读取t_knowledge_tree表中codeName和answer字段
	var id, name string
	errTables := db.QueryRow("SELECT * FROM user WHERE id>", quitionID).Scan(&id, &name)
	if errTables != nil {
		//TODO，这里只是打印了一下，并没有做异常处理
		fmt.Println("funReadSql SELECT t_knowledge_tree is error", errTables)
	}
	fmt.Printf("funReadSql codeName:%s, answer:%s\n", id, name)

	//关闭数据库
	db.Close()
	return id, name
}
