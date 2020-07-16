package main

import (
	"fmt"
	"github.com/Limard/websocket"
	"net/rpc"
	"os"
	"time"
)

func main() {
	type TESTBuf struct {
		Buf []byte
	}

	conn, _, e := websocket.DefaultDialer.Dial("ws://127.0.0.1:7878", nil)
	if e != nil {
		panic(e)
	}

	client := rpc.NewClient(conn)

	f, e := os.Open(`C:\Users\ThinkPad\Downloads\elasticsearch-6.3.2.zip`)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	res := &struct{}{}
	//buf := make([]byte, 4 * 1024 - 200)
	buf := make([]byte, 4 * 1024)
	for i := 0; i < 2; i++ {
		n, e := f.Read(buf)
		if e != nil {
			break
		}

		fmt.Println("send buffer ", n)
		e = client.Call("TESTRPC.BUF", &TESTBuf{Buf: buf[:n]}, res)
		if e != nil {
			panic(e)
		}
	}

	time.Sleep(time.Second)
}
