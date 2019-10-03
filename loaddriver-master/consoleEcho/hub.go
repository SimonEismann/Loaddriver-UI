package consoleEcho

import (
	"loaddriver-master/shared"
	"net/http"

	"loaddriver-master/logger"
)

type hub struct {
	logger      *logger.StandardLogger
	forward     chan []byte
	join        chan *consoleSubscriber
	leave       chan *consoleSubscriber
	subscribers map[*consoleSubscriber]bool
}

func (h *hub) Run() {
	for {
		select {
		case sub := <-h.join:
			h.subscribers[sub] = true
		case sub := <-h.leave:
			h.logger.WithField("comp", "hub").WithField("func", "write").Info("Client unsubscribed")
			delete(h.subscribers, sub)
			close(sub.send)
		case msg := <-h.forward:
			for sub := range h.subscribers {
				sub.send <- msg
			}
		}
	}
}

func (h *hub) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := shared.NewUpgrader().Upgrade(w, req, nil)
	if err != nil {
		h.logger.WithField("comp", "hub").WithField("func", "ServeHTTP").WithError(err).Error("Error during upgrade")
		return
	}
	client := &consoleSubscriber{
		logger: h.logger,
		socket: socket,
		send:   make(chan []byte, shared.MessageBufferSize),
		hub:    h,
	}
	h.join <- client
	defer func() { h.leave <- client }()
	h.logger.WithField("comp", "hub").WithField("func", "ServeHTTP").Info("New subscriber registered")
	go client.write()
	client.read()
}

func NewHub(logger *logger.StandardLogger, forward chan []byte) *hub {
	return &hub{
		logger:      logger,
		forward:     forward,
		join:        make(chan *consoleSubscriber),
		leave:       make(chan *consoleSubscriber),
		subscribers: make(map[*consoleSubscriber]bool),
	}
}
