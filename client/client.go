package main
//
//import (
//	"log"
//	"net"
//	"time"
//)
//
///**
//  @author: ZH
//  @since: 2023/11/2
//  @desc: //TODO
//**/
//
//func main() {
//	log.Println("Client Test ... start")
//	time.Sleep(3 * time.Second)
//
//	conn, err := net.Dial("tcp", "127.0.0.1:7777")
//	if err != nil {
//		log.Println("Client start err:", err)
//		return
//	}
//	for {
//		_, err := conn.Write([]byte("haha"))
//		if err != nil {
//			log.Println("write error err:", err)
//			return
//		}
//		buf := make([]byte, 512)
//		cnt, err := conn.Read(buf)
//		if err != nil {
//			log.Println("read buf error:", err)
//			return
//		}
//		log.Printf("server call back:%s,cnt:%d  \n", buf, cnt)
//		time.Sleep(1 * time.Second)
//	}
//}
