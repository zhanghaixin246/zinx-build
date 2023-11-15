package znet

import (
	"log"
	"net"
	"testing"
	"time"
)

/**
  @author: ZH
  @since: 2023/11/2
  @desc: //TODO
**/

func ClientTest() {
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Println("client start err: ", err)
		return
	}
	for {
		_, err := conn.Write([]byte("hello ZINX"))
		if err != nil {
			log.Println("write error:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Println("read buf error:", err)
			return
		}
		log.Printf("server call back :%s,cnt=%d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer()
	go ClientTest()
	s.Serve()
}
