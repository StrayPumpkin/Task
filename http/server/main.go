package main

import (
	"fmt"
	"net/http"
)

// helloHandler 是处理HTTP请求的函数。第一个参数用于向客户端发送响应；第二个参数包含了客户端请求的所有信息。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // 检查请求方法是否为GET
		fmt.Fprint(w, "Hello, Client!") // 向客户端写入响应内容
	}
}

func main() {
	//将根路径"/"与helloHandler关联起来。
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server at :8090")
	//启动HTTP服务器
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println(err)
	}
}