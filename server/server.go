package main

import (
	"log"
	"net"
	"reflect"
	"strings"
	"time"
)

// 1. 限流(令牌桶限流)
// 2. 鉴权
// 3. 路由
var (
	ch        chan int = make(chan int, 5)
	WhiteList []string = []string{"127.0.0.1", "192.168.1.10"}
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("net.Listen error", err.Error())
		return
	}
	defer l.Close()
	defer close(ch)

	log.Println("tcp listen on 8080")

	go addToken()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("l.Accept() error", err.Error())
			return
		}

		// 1. 限流
		<-ch
		log.Println("has token")

		// 2. 鉴权
		if auth(conn) {
			// 3. 路由
			go route(conn)
		}
	}
}

func auth(newConn net.Conn) bool {
	remoteIp := strings.Split(newConn.RemoteAddr().String(), ":")[0]
	log.Println("remote ip: ", remoteIp)

	if !sliceContains(remoteIp) {
		if _, err := newConn.Write([]byte("No permission")); err != nil {
			log.Println("write No permission error", err.Error())
			return false
		}

		return false
	}

	return true
}

func sliceContains(ip string) bool {
	for _, v := range WhiteList {
		if v == ip {
			return true
		}
	}

	return false
}

// 添加限流令牌,1s一次访问
func addToken() {
	for {
		time.Sleep(time.Second * 1)
		ch <- 1
	}
}

// 客户端断开处理
func onDisConnect() {
	log.Println("client close conn")
}

func route(newConn net.Conn) {
	buf := make([]byte, 2048)
	_, err := newConn.Read(buf)

	if err != nil {
		log.Println("read error", err.Error())
		if err.Error() == "EOF" {
			onDisConnect()
		}
		return
	}

	bufstring := string(buf)
	log.Println("server read msg:", bufstring)

	funcName := strings.Split(bufstring, "/")[0]
	arg := strings.Split(bufstring, "/")[1]

	log.Println("funcName, arg:", funcName, arg)

	refectValueArg := reflect.ValueOf(arg)
	typeHandler := &Handler{}

	in := []reflect.Value{refectValueArg}
	returnValues := reflect.ValueOf(typeHandler).MethodByName(funcName).Call(in)
	returnValue := returnValues[0].String()

	log.Println("returnValues:", returnValues[0].String())

	if len(returnValue) > 0 {
		if _, err = newConn.Write([]byte(returnValue)); err != nil {
			log.Println("write error", err.Error())
		}
	}
}

type Handler struct {
}

func (h *Handler) GetName(arg string) string {
	return arg + " " + "world"
}
