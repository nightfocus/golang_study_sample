package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func init() {
	quitSemaphore = make(chan bool)
}

// var t1 int32

// gl 用来控制允许的最大并发数
// cnt 是连接TCP成功后，发送报文的个数
func echoClientLimited(addr string, cnt int, ngl *GoLimit) {
	defer ngl.Done()

	echoClient(addr, cnt)
}

func echoClient(addr string, cnt int) {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", addr)

	time.Sleep(time.Second) // wait for tcp server is ready.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Println("tcp is connected! ", addr, conn.LocalAddr().String())

	go onMessageRecived(conn)

	totalsend := 0
	for i := 0; i < cnt; i++ {
		b := make([]byte, 1000+i)
		b[0] = 0xC2
		s, _ := conn.Write(b)
		totalsend += s
		time.Sleep(1 * time.Millisecond)

		fmt.Println("cnt: send s", i, s)

		/*
			if t1 > 3000 {
				fmt.Println("long time no recv!!!! ", t1)
				break
			}
			t1++
		*/
	}

	select {
	case <-quitSemaphore:
		fmt.Println("need to close self: bytes: ", totalsend, conn.LocalAddr().String())
	}

}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	totalrecv := 0
	for {
		len, err := reader.Read(buf)
		if err != nil {
			fmt.Println("read error! all recv bytes: ", totalrecv, err)
			quitSemaphore <- true
			break
		}
		totalrecv += len
		// fmt.Println("recv len: ", len)
		// t1 = 0
		// time.Sleep(time.Second)
		// b := []byte(msg)
		// conn.Write(b)
	}
}
