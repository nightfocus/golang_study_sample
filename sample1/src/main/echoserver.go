package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func echoServer() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")

	tcpListener, _ := net.ListenTCP("tcp4", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()

	// 设置3秒的读取超时
	// 注意：这个函数语义是设定一个绝对时间值。也就是在reader.Read(..)超时后，如果不重新设置，
	// 那么再次调用reader.Read(..) 会一直超时.
	if err := conn.SetReadDeadline(time.Now().Add(time.Second * 3)); err != nil {
		fmt.Println("err: ", err)
		return
	}

	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	for {
		len, err := reader.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}

		fmt.Println(string(buf))
		conn.Write(buf[0:len])
		// _ = len
	}
}
