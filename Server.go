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

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		log.Println("err:", err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (hzr *HelloZinxRouter) Handle(request ziface.IRequest) {
	log.Println("call HelloZinxRouter Handle")
	log.Println("receive from client: msgId=", request.GetMsgId(), " data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("hello Zinx Router v0.6"))
	if err != nil {
		log.Println("err:", err)
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
