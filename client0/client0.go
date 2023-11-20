package main

import (
	"io"
	"log"
	"net"
	"time"
	"zinx-build/znet"
)

/*
*

	@author: ZH
	@since: 2023/11/17
	@desc: //TODO

*
*/
func main() {

	log.Println("Client Test ... start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Println("client start err,exit!")
		return
	}
	for {
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx V0.6 client0 test message")))
		_, err := conn.Write(msg)
		if err != nil {
			log.Println("write error:", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			log.Println("read head error")
			break
		}

		msgHead, err := dp.UnPack(headData)
		if err != nil {
			log.Println("server UnPack err:", err)
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
			log.Println("==> receive msg id=", msg.Id, " len=", msg.DataLen, " data=", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}
}
