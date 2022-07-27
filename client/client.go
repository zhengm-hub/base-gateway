package main

import (
	"log"
	"net"
)

const (
	addr = "127.0.0.1:8080"
)

func main() {
	for {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Println("net.Listen error", err.Error())
			return
		}
		defer conn.Close()

		if _, err = conn.Write([]byte("GetName/hello")); err != nil {
			log.Println("write error", err.Error())
			return
		}

		buf := make([]byte, 2048)
		_, err = conn.Read(buf)
		if err != nil {
			log.Println("read error", err.Error())
			return
		}

		log.Println("client read msg:", string(buf))
	}
}
