package socketio

import (
	"bytes"
	"github.com/gorilla/websocket"
	"io"
	"time"
)

func init() {
	DefaultTransports.RegisterTransport("websocket")
}

type webSocket struct {
	session *Session
	timeout time.Duration
	conn    *websocket.Conn
}

func newWebSocket(session *Session) *webSocket {
	ret := &webSocket{
		session: session,
		timeout: session.heartbeatTimeout / 10,
	}
	session.transport = ret
	return ret
}

func (ws *webSocket) Send(data []byte) error {
	return ws.conn.WriteMessage(websocket.TextMessage, data)
}

func (ws *webSocket) Read() (io.Reader, error) {
	_, message, err := ws.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(message)
	return reader, nil
}
