package ziface

import "net"

/**
  @author: ZH
  @since: 2023/11/2
  @desc: //TODO
**/

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}

type HandFunc func(*net.TCPConn, []byte, int) error
