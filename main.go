package main

import (
	"io"
	"log"
	"net"
	"time"
)

//$ go build gopl.io/ch8/clock1
//$ ./clock1 &
//$ nc localhost 8000
func main() {
	//net.Listen 对象 监听来自8001端口的连接
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//listener.Accept方法会直接阻塞，直到新的连接被创建 返回 net.Conn对象来表示这个连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		//处理一个完整的客户端连接
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	//死循环会一直执行，直到写入失败
	//最可能的原因是客户端主动断开连接
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
