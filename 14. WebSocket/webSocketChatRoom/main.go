package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
	roomIdToClients = map[string][]*websocket.Conn{}
)


func wsChatRoom(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	roomId := c.Param("roomId")
	username := c.Param("username")
	
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error("upgrade err:", err)
		return err
	}
	defer ws.Close()

	clients := roomIdToClients[roomId]
	clients = append(clients, ws)
	roomIdToClients[roomId] = clients
	fmt.Println("Add new client to room:", roomId, roomIdToClients)

	for {
		// Read new messages
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
			return err
		} else {
			fmt.Println("New msg:", string(msg))
			for _, conn := range roomIdToClients[roomId] {
				fmt.Println("loop over connections:", &conn, &ws)
				if conn != ws {
					fmt.Println("Send new message to existing client", string(msg), roomIdToClients)
					err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", username, msg)))
					if err != nil {
						c.Logger().Error(err)
						return err
					}
				}
			} 
		}
	}
	return c.String(http.StatusOK, roomId)
}

func main() {
    e := echo.New()
	e.GET("/ws/chat/:roomId/user/:username", wsChatRoom)
	e.Logger.Fatal(e.Start(":8080"))
}