package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"musicBiuBiu/apis"
	"myGo/fantasticmao/websocket/api"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		//跨域访问
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/biubiu", apis.Biubiu)
	router.GET("/spider/:songId", apis.Spider)
	router.GET("/getComment/:musicName", apis.GetCommentsByMusicName)
	return router
}
func main() {

	router := initRouter()
	router.Run(":8080")
}

//ws访问
func WsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *api.Connection
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = api.InitConnection(wsConn); err != nil {
		goto ERROR
	}

	go func() {
		if err = conn.WriteMessage([]byte("heartBeat")); err != nil {
			goto ERROR
		}
	ERROR:
		conn.Close()
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERROR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERROR
		}
	}
ERROR:
	conn.Close()

}
