package ziface

/**
  @author: ZH
  @since: 2023/11/3
  @desc: //TODO
**/

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
