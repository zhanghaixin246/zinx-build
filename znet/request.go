package znet

import "zinx-build/ziface"

/**
  @author: ZH
  @since: 2023/11/3
  @desc: //TODO
**/

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32{
	return r.msg.GetMsgId()
}
