package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("client start...")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "localhost:8999")
	if err != nil {
		log.Printf("client connect error: %v\n", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello gnet v0.2"))
		if err != nil {
			log.Printf("client write error: %v\n", err)
			return
		}

		// 写数据
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Printf("client read error: %v\n", err)
			return
		}

		log.Printf("client read: %s, cnt: %d\n", buf[:cnt], cnt)

		// block cpu for seconds
		time.Sleep(1 * time.Second)
	}
}
