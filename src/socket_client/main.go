package socket_client

import (
	"fmt"
	"net"
)

func NewClient(addr string) bool {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务器失败")
	}
	defer conn.Close()
	fmt.Println("连接服务成功")
	conn.Write([]byte("客户端请求连接服务器"))
	data := make([]byte, 255)
	conn.Read(data)
	fmt.Println(string(data))
	return true
}
