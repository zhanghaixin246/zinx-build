package ziface

/**
  @author: ZH
  @since: 2023/11/2
  @desc: //TODO
**/

type IServer interface {
	Start()
	Stop()
	Serve()

	AddRouter(router IRouter)
}
