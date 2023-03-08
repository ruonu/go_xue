package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello") // 服务端打印输出
	fmt.Fprintf(w, "hello GoLangWEB")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login") // 服务端打印输出
	fmt.Fprintf(w, "login...")
}

func main() {
	http.HandleFunc("/", Hello)      // 匹配根路由
	http.HandleFunc("/login", login) // 匹配根路由/login
	err := http.ListenAndServe("0.0.0.0:18001", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
