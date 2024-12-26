package main

import (
	"server/database"
	"server/views"
	"server/views/api"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var manager = NewWebSocketManager()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func createWebsocket(context *gin.Context) {
	socket, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
	manager.AddClient(socket)
	go manager.Clients[socket].Loop()
}

func setupViews(router *gin.Engine, db *database.DB) {
	router.GET("/", views.Index)
	router.GET("/websocket", createWebsocket)
	router.POST("/api/register", func(c *gin.Context) {
		api.Register(c, db)
	})
	router.LoadHTMLGlob("./templates/*")
}

func main() {
	manager.db = database.SetupDB()
	router := gin.Default()
	setupViews(router, manager.db)
	router.Run(":80")
}
