package znet

import (
	"errors"
	"fmt"
	"io"
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
		//buf := make([]byte, 512)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("receive buf err:", err)
		//	c.ExitBuffChan <- true
		//	continue
		//}

		//buf := make([]byte,utils.GlobalObject.MaxPacketSize)
		//_,err := c.Conn.Read(buf)
		//if err != nil {
		//	log.Println("receive buf err:",err)
		//	c.ExitBuffChan <- true
		//	continue
		//}
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			log.Println("read msg head error:", err)
			c.ExitBuffChan <- true
			continue
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			log.Println("unpack error", err)
			c.ExitBuffChan <- true
			continue
		}
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				log.Println("read msg data error:", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		req := Request{
			conn: c,
			msg:  msg,
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		log.Println("pack error msg id=", msgId)
		return errors.New("pack error msg")
	}

	if _, err := c.Conn.Write(msg); err != nil {
		log.Println("write msg id ", msgId, " error")
		c.ExitBuffChan <- true
		return errors.New("conn write error " + err.Error())
	}
	return nil
}
