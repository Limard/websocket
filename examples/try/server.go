package main

import (
	"fmt"
	"github.com/Limard/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("running...")
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(":7878", nil))
}

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new...")
	c, err :=  upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	c.SetPingHandler(func(appData string) error {
		fmt.Println("ping handler")
		return c.WriteControl(websocket.PongMessage, nil, time.Now().Add(time.Second))
	})

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}