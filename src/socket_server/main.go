package socket_server

import (
	"fmt"
	"net"
)

func NewServer(addr string) interface{} {
	socket, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("服务端创建socket失败+v%", err)
	}
	defer socket.Close()
	conn, err := socket.Accept()
	if err != nil {
		fmt.Println("服务端socket接受请求失败")
	}
	data := make([]byte, 255)
	msgLength, err := conn.Read(data)
	if err != nil {
		fmt.Println("服务端读取数据失败")
	}
	fmt.Println("服务端读取数据v%", msgLength)
	conn.Write([]byte("服务端:收到了"))

	return socket
}
