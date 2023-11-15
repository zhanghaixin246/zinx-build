package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx-build/ziface"
)

/**
  @author: ZH
  @since: 2023/11/14
  @desc: //TODO
**/

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	Version   string

	MaxPacketSize uint32
	MaxConn       int
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload(){
	data,err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data,&GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init(){
	GlobalObject = &GlobalObj{
		//TcpServer:     nil,
		Host:          "0.0.0.0",
		TcpPort:       7777,
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		MaxPacketSize: 12000,
		MaxConn:       4096,
	}
	GlobalObject.Reload()
}