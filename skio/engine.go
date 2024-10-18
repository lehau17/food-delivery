package skio

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/jwtprovider/jwt"
	userstorage "github.com/lehau17/food_delivery/modules/user/storage"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx appcontext.AppContect, engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewRtEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *rtEngine) SaveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}
	engine.locker.Unlock()

}

func (engine *rtEngine) GetAppSocket(userId int) []AppSocket {
	engine.locker.Lock()
	defer engine.locker.Unlock()
	return engine.storage[userId]
}

func (engine *rtEngine) RemoveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()
	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}

		}
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket
	if scks, ok := engine.storage[userId]; ok {
		return scks
	}
	return sockets
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.storage[userId]
	for _, socket := range sockets {
		socket.Emit(key, data)
	}
	return nil
}

func (engine *rtEngine) Run(appCtx appcontext.AppContect, r *gin.Engine) error {
	server := socketio.NewServer(&engineio.Options{Transports: []transport.Transport{websocket.Default}})
	// log.Println(server.)
	engine.server = server
	server.OnConnect("", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:>>>", s.ID(), "IP :>>>", s.RemoteAddr(), s.ID())
		return nil
	})
	server.OnError("/", func(c socketio.Conn, err error) {
		fmt.Println("meet error", err)
	})
	server.OnDisconnect("/", func(c socketio.Conn, s string) {
		fmt.Println("reason :>>>", s)

	})
	server.OnEvent("/", "authentication", func(s socketio.Conn, token string) {
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStore(db, appCtx.GetRedis())
		tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}
		user, err := store.Find(context.Background(), map[string]interface{}{"id": payload.Uid})
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}
		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("authentication error"))
			s.Close()
			return
		}
		user.Mask(false)
		appSck := NewAppSocket(s, user)
		engine.SaveAppSocket(payload.Uid, appSck)
		s.Emit("authenticated", user)

	})
	go server.Serve()
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	log.Println("SKIO CONNECT")
	return nil
}
