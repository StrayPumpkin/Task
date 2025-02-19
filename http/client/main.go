package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 发起一个GET请求到服务器。
	resp, err := http.Get("http://localhost:8090/")
	if err != nil { // 如果请求过程中发生错误，则打印错误并退出。
		fmt.Println("Error fetching URL:", err)
		return
	}
	//在函数返回之前关闭，释放资源
	defer resp.Body.Close()

	// 读取服务器响应的数据。
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { 
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Message from server:", string(body))
}