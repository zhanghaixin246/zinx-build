package main

import (
	"io"
	"log"
	"net"
	"zinx-build/znet"
)

/**
  @author: ZH
  @since: 2023/11/15
  @desc: //TODO
**/

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("server accept err:", err)
		}
		go func(conn net.Conn) {
			dp := znet.NewDataPack()
			for {
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					log.Println("read head err:", err)
					break
				}
				msgHead, err := dp.UnPack(headData)
				if err != nil {
					log.Println("server unpack err:", err)
					return
				}

				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						log.Println("server unpack data err:", err)
						return
					}
					log.Println("==> Recvive Msg:ID=", msg.Id, " len=", msg.DataLen, " data=", string(msg.Data))
				}
			}
		}(conn)
	}
}
