package skio

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/lehau17/food_delivery/common"
)

type AppSocket interface {
	socketio.Conn
	common.Requester
}

type appSocket struct {
	socketio.Conn
	common.Requester
}

func NewAppSocket(conn socketio.Conn, requester common.Requester) *appSocket {
	return &appSocket{Conn: conn, Requester: requester}
}
