package main

import (
	"log"
	"zinx-build/ziface"
	"zinx-build/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	log.Println("call PingRouter Handle")
	log.Println("receive from client:msgId=", request.GetMsgId(), " data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		log.Println("err:", err)
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
