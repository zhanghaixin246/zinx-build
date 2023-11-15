package ziface

/**
  @author: ZH
  @since: 2023/11/15
  @desc: //TODO
**/

type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
