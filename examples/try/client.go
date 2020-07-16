package main

import (
	"fmt"
	"github.com/Limard/websocket"
	"log"
	"time"
)

func main() {
	fmt.Println("dial")
	url_ := `ws://127.0.0.1:7878/`
	c, _, e := websocket.DefaultDialer.Dial(url_, nil)
	if e != nil {
		panic(e)
	}

	pongTimeout := time.NewTimer(30 * time.Second)

	c.SetPongHandler(func(appData string) error {
		fmt.Println("pong handler")
		pongTimeout.Reset(30 * time.Second)
		return nil
	})

	done := make(chan struct{})

	// read
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("ping")
			e = c.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second))
			if e != nil {
				fmt.Println("e:", e)
				return
			}
		case <-done:
			fmt.Println("done")
			return
		case <-pongTimeout.C:
			fmt.Println("Pong Timeout")
			return
		}
	}
}
