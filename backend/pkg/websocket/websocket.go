package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, nil
}

//func Reader(conn *websocket.Conn) {
//	for {
//		//	read in a message
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		//	print out a message
//		fmt.Println(string(p))
//
//		if err = conn.WriteMessage(messageType, p); err != nil {
//			log.Println(err)
//			return
//		}
//	}
//}
//
//func Writer(conn *websocket.Conn) {
//	for {
//		fmt.Println("Sending")
//		messageType, r, err := conn.NextReader()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		w, ErrW := conn.NextWriter(messageType)
//		if ErrW != nil {
//			log.Println(ErrW)
//			return
//		}
//		if _, err = io.Copy(w, r); err != nil {
//			log.Println(err)
//			return
//		}
//		if err = w.Close(); err != nil {
//			log.Println(err)
//			return
//		}
//	}
//}
