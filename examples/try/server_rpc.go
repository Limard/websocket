package main

import (
	"fmt"
	"github.com/Limard/websocket"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	fmt.Println("running...")

	rpc.DefaultServer.RegisterName("TESTRPC", new(TESTRPC))

	upgrader := websocket.Upgrader{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("new...")
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		rpc.DefaultServer.ServeConn(c)
		fmt.Println("exit")
	})
	log.Fatal(http.ListenAndServe(":7878", nil))
}

type TESTRPC struct{}

type TESTBuf struct {
	Buf []byte
}

func (TESTRPC) BUF(req *TESTBuf, res *struct{}) (e error) {
	fmt.Println("len:", len(req.Buf))
	return
}
