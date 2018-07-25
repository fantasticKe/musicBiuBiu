package push

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	WsConn      *websocket.Conn
	Rc          chan []byte
	Wc          chan []byte
	CloseChan   chan byte
	mutex       sync.Mutex
	IsCloseChan bool
}

//初始化连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		WsConn:    wsConn,
		Rc:        make(chan []byte, 1000),
		Wc:        make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}

	go conn.readLoop()

	go conn.writeLoop()
	return
}

//
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.Rc:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.Wc <- data:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	conn.WsConn.Close()
	conn.mutex.Lock()
	if !conn.IsCloseChan {
		close(conn.CloseChan)
		conn.IsCloseChan = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.WsConn.ReadMessage(); err != nil {
			goto ERROR
		}
		select {
		case conn.Rc <- data:
		case <-conn.CloseChan:
			goto ERROR
		}

	}

ERROR:
	conn.WsConn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.Wc:
		case <-conn.CloseChan:
			goto ERROR
		}
		if err = conn.WsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERROR
		}
	}
ERROR:
	conn.WsConn.Close()
}
