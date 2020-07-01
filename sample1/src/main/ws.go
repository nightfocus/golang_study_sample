package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域访问
		// CheckOrigin 是Upgrader的成员，类型是函数
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

	// 收到http请求(upgrade),完成websocket协议转换
	// 在应答的header中放上upgrade:websoket
	var (
		conn *websocket.Conn
		err  error
		//msgType int
		data []byte
	)

	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		// 报错了，直接返回底层的websocket链接就会终断掉
		return
	}

	// 得到了websocket.Conn长连接的对象，实现数据的收发
	for {
		// Text(json), Binary
		//if _, data, err = conn.ReadMessage(); err != nil {
		if _, data, err = conn.ReadMessage(); err != nil {
			// 报错关闭websocket
			goto ERR
		}
		// 发送数据，判断返回值是否报错
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			// 报错了
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

/*
func main() {
	//http://localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/hello", helloHandler)
	//服务端启动
	http.ListenAndServe("0.0.0.0:7777", nil)
}
*/
