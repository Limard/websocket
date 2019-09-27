package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, e := net.Dial(`tcp`, `192.168.137.174:28080`)
	if e != nil {
		panic(e)
	}

	go func() {
		buf := make([]byte, 512)
		for {
			n, e := conn.Read(buf)
			fmt.Println("read", n, e)
			if e != nil {
				break
			}
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	var n int
	for {
		n, e = conn.Write([]byte("123456"))
		fmt.Println("ping", n, e)
		if e != nil {
			break
		}
		<-ticker.C
	}
}
