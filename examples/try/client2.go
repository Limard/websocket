package main

import (
	"fmt"
	"github.com/Limard/websocket"
	"os"
	"time"
)

func main() {
	//conn, e := net.Dial(`tcp`, `127.0.0.1:7878`)
	conn, _, e := websocket.DefaultDialer.Dial("ws://127.0.0.1:7878", nil)
	if e != nil {
		panic(e)
	}

	go func() {
		buf := make([]byte, 512)
		for {
			n, e := conn.Read(buf)
			fmt.Println("read:", n, e)
			fmt.Println("read:", string(buf[:n]))
			if e != nil {
				break
			}
		}
	}()

	//ticker := time.NewTicker(5 * time.Second)
	//var n int
	//for {
	//	n, e = conn.Write([]byte("123456"))
	//	fmt.Println("ping", n, e)
	//	if e != nil {
	//		break
	//	}
	//	<-ticker.C
	//}

	f, e := os.Open(`C:\Users\ThinkPad\Downloads\elasticsearch-6.3.2.zip`)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	buf := make([]byte, 64*1024)
	for i := 0; i < 10; i++ {
		n, e := f.Read(buf)
		if e != nil {
			break
		}

		n, e = conn.Write(buf[:n])
		if e != nil {
			panic(e)
		}
		fmt.Println("write:", n)
	}

	time.Sleep(time.Second)
}
