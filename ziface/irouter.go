package ziface

/**
  @author: ZH
  @since: 2023/11/3
  @desc: //TODO
**/

type IRouter interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
