package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
	"time"
)

func main(){
	fmt.Println("客户端启动")
	//调用Dial函数: 参数需要指定tcp协议，需要指定服务器端的IP+port
	conn,err := net.Dial("tcp","192.168.171.10:8888")
	if err !=nil{//连接失败
		fmt.Println("客户端连接失败: err:",err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功,conn:",conn)
	//通过客户端发送数据，然后退出
	reader := bufio.NewReader(os.Stdin)//os.Stdin代表终端标准输出

	//从终端读取一行用户输入的信息
	fmt.Println("请输入数据：")
	str,err := reader.ReadString('\n')
	if err !=nil{
		fmt.Println("终端输入失败: err:",err)
		return
	}
	//将str数据发送给服务器
	n,err := conn.Write([]byte(str))
	if err !=nil{
		fmt.Println("发送数据失败: err:",err)
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节数据,并退出\n",n)

	//保持连接打开一段时间以确保服务器有足够时间处理数据
	time.Sleep(1 * time.Second)
	fmt.Println("客户端即将退出。。。")
}