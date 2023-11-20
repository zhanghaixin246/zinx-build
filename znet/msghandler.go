package znet

import (
	"log"
	"strconv"
	"zinx-build/utils"
	"zinx-build/ziface"
)

/**
  @author: ZH
  @since: 2023/11/15
  @desc: //TODO
**/

type MsgHandle struct {
	Apis           map[uint32]ziface.IRouter
	WorkerPoolSize uint32
	TaskQueue      []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

func (mh *MsgHandle) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	log.Println("worker id=", workerId, " is started.")
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	workerId := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	log.Println("Add ConnID:", request.GetConnection().GetConnID(), " request msgId=", request.GetMsgId(), " to workerId=", workerId)
	mh.TaskQueue[workerId] <- request
}
