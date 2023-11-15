package main

import (
	"log"
	"zinx-build/ziface"
	"zinx-build/znet"
)

/**
  @author: ZH
  @since: 2023/11/14
  @desc: //TODO
**/

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	log.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PreHandle ...\n"))
	if err != nil {
		log.Println("call back ping error" + err.Error())
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	log.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Handle ... \n"))
	if err != nil {
		log.Println("call back ping  handle error" + err.Error())
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	log.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PostHandle ...\n"))
	if err != nil {
		log.Println("call back ping  postHandle error " + err.Error())
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()

}
