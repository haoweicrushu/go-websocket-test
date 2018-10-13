package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)


var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func chatroom(c *gin.Context) {

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "foo" {
			message = []byte("bar")
		}

		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	bindAddress := "localhost:8888"
	r := gin.Default()
	r.GET("/chatroom", chatroom)
	r.Run(bindAddress)
}
