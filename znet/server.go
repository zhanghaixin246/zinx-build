package znet

import (
	"errors"
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

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("CallBackToClient error:" + err.Error())
	}
	return nil
}

func (s *Server) Start() {
	log.Printf("[START] Server listenner at IP:%s,Port %d,is starting \n", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Println("resolve tcp addr err:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen:", s.IPVersion, " err:", err)
			return
		}
		log.Println("Start Zinx server ", s.Name, " success,now listening...")
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("Accept err:", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	log.Println("[STOP] Zinx server , name ", s.Name)
	// todo server stop
}

func (s *Server) Serve() {
	s.Start()
	// todo server serve
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
		Router:    nil,
	}
	return s
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	log.Println("AddRouter success ! ")
}
