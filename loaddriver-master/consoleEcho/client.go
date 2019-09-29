package consoleEcho

import (
	"loaddriver-master/logger"

	"github.com/gorilla/websocket"
)

type consoleSubscriber struct {
	logger *logger.StandardLogger
	socket *websocket.Conn
	send   chan []byte
	hub    *hub
}

func (c *consoleSubscriber) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
