package znet

import (
	"fmt"
	"log"
	"net"
	"zinx-build/ziface"
)

/**
  @author: ZH
  @since: 2023/11/2
  @desc: //TODO
**/

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	//handleAPI    ziface.HandFunc
	Router       ziface.IRouter
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI:    callbackApi,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	log.Println("Reader Goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " Conn reader exit!")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err:", err)
			c.ExitBuffChan <- true
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}

}

func (c *Connection) Start() {
	go c.StartReader()
	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.Conn.Close()
	c.ExitBuffChan <- true
	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
