package main

import (
	"fmt"
	"net"
	"os"

	// "time"

	. "ctoola"
)

const (
	MAX_CONN_NUM = 3
)

//echo server Goroutine
func EchoFunc(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 32)
	for {
		readLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Printf("read len: %d, %v\n", readLen, buf)
		//send reply
		writeLen, err := conn.Write(buf[0:readLen])
		if err != nil {
			//fmt.Printf("Error send reply:", err.Error())
			return
		}
		fmt.Printf("write len: %d\n", writeLen)
	}
}

//initial listener and run
func netserver() {
	listener, err := net.Listen("tcp", "0.0.0.0:9088")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Network server is running ...\n")

	var cur_conn_num int = 0
	conn_chan := make(chan net.Conn)
	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
			fmt.Printf("cur connect num: %v\n", cur_conn_num)
		}
	}()

	/*
		go func() {
			for _ = range time.Tick(1e8) {
				fmt.Printf("cur conn num: %f\n", cur_conn_num)
			}
		}()
	*/

	// 直接按最大允许的连接数，开启对应的协程。每个协程用range conn_chan 去等待新的连接.
	// 然后进行echo. 以达到真正并发的处理.
	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range conn_chan {
				fmt.Printf("goroutine %d: recv conn.\n", GetGoroutineID())
				ch_conn_change <- 1
				EchoFunc(conn)
				ch_conn_change <- -1
			}
		}()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accept:", err.Error())
			return
		}
		conn_chan <- conn
	}
}
