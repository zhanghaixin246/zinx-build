package ziface

/**
  @author: ZH
  @since: 2023/11/15
  @desc: //TODO
**/

type IMsgHandle interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}
