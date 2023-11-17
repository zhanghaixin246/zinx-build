package znet

import (
	"log"
	"strconv"
	"zinx-build/ziface"
)

/**
  @author: ZH
  @since: 2023/11/15
  @desc: //TODO
**/

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		log.Println("api msgId=", request.GetMsgId(), " not found")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PreHandle(request)
}

func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api,msgId=" + strconv.Itoa(int(msgId)))
	}
	mh.Apis[msgId] = router
	log.Println("add api msgId=", msgId)
}
