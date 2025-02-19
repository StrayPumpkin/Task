package main

import(
	"fmt"
	"net"
)

func process(conn net.Conn){
	//连接用完一定要关闭
	defer conn.Close()
	fmt.Printf("开始处理来自%v的连接\n",conn.RemoteAddr().String())
	
	//创建一个切片，准备将读取的数据放入切片
	buf := make([]byte,1024)
	for{
		//从conn连接中读取数据
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("数据读取失败：",err)
			return
		}
		//将读取内容在服务器端输出
		fmt.Printf("接收到客户端 %v 的数据: %s\n",conn.RemoteAddr().String(),string(buf[0:n]))
	}
}

func main(){
	fmt.Println("服务器端启动。")
	//监听：指定服务器端TCP协议，服务器端的IP+port
	listen, err := net.Listen("tcp","192.168.171.10:8888")
	if err != nil{//监听失败
		fmt.Println("net.Listen err:", err)
		return
	}
	//循环等待客户端的连接
	for {
		conn,err2 := listen.Accept()
		if err2 != nil {//客户端等待失败
			fmt.Println("客户端等待失败,err2:",err2)
		}else{//连接成功
			fmt.Printf("等待连接成功,conn = %v,接收到客户端信息：%v \n",conn,conn.RemoteAddr().String())
		}
		//准备一个协程，协程处理客户端服务请求
		go process(conn)
		//不同客户端的请求连接conn不一样
	} 
}	