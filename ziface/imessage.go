package ziface

/**
  @author: ZH
  @since: 2023/11/14
  @desc: //TODO
**/

type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetData([]byte)

	SetDataLen(uint32)
}
