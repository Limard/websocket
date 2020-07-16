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

	upgrader := websocket.Upgrader{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("new...")
		c, err := upgrader.Upgrade(w, r, nil)
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
			log.Printf("recv: %d", len(message))
			err = c.WriteMessage(mt, []byte("get data"))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	log.Fatal(http.ListenAndServe(":7878", nil))
}
