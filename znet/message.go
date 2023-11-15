package znet

/**
  @author: ZH
  @since: 2023/11/15
  @desc: //TODO
**/

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMsgPackage(id uint32,data []byte)*Message{
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (msg *Message) GetDataLen()uint32{
	return msg.DataLen
}

func (msg *Message) GetMsgId() uint32{
	return msg.Id
}

func (msg *Message) GetData() []byte{
	return msg.Data
}

func (msg *Message) SetDataLen(len uint32){
	msg.DataLen = len
}

func (msg *Message) SetMsgId(msgId uint32){
	msg.Id = msgId
}

func (msg *Message) SetData(data []byte){
	msg.Data = data
}