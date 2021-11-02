package skio

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	answerbus2 "VoteSth-socketgo/modules/answer/bus"
	answerstorage "VoteSth-socketgo/modules/answer/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"sync"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int,key string,data interface{}) error
	Run(ctx appctx.AppContext, engine *gin.Engine) error
}

type rtEngine struct{
	server *socketio.Server
	storage map[int][]AppSocket
	locker *sync.RWMutex
}

func NewEngine() *rtEngine{
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker: new(sync.RWMutex),
	}
}
func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	//appSck.Join("order-{ordID}")

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
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
func (engine *rtEngine ) UserSockets(userId int)[]AppSocket{
	var sk []AppSocket
	if scks,ok := engine.storage[userId];ok{
		return scks
	}
	return sk
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.Emit(key, data)
	}

	return nil
}

func (engine *rtEngine) Run(appCtx appctx.AppContext, r *gin.Engine) error {
	server := socketio.NewServer(nil)
	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr(), s.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	// Setup

	server.OnEvent("/", "getlistanswer", func(s socketio.Conn, msg string) {
		db := appCtx.GetMainDbConnection()
		store := answerstorage.NewSQLStore(db)
		bus:=answerbus2.NewListAnswerBus(store)
		uid, _ := common.FromBase64(msg)
		fmt.Println(uid.GetLocalID())
		data,err := bus.GetAnswerByQuestionId(context.Background(), int(uid.GetLocalID()))
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		for i := range data{
			data[i].Mask()
		}
		appSck := NewAppSocket(s)
		engine.saveAppSocket(int(uid.GetLocalID()), appSck)

		s.Emit("listanswer", data)

	})
	server.OnEvent("/", "vote", func(s socketio.Conn, msg string) {
		db := appCtx.GetMainDbConnection()
		//Vote
		store := answerstorage.NewSQLStore(db)
		bus:=answerbus2.NewVoteAnswerBus(store)
		uid, _ := common.FromBase64(msg)
		bus.VoteAnswer(context.Background(),int(uid.GetLocalID()))
		//LoadData
		bus2 := answerbus2.NewListAnswerBus(store)
		data,err := bus2.GetAnswerByQuestionId(context.Background(),int(uid.GetLocalID()))
		if err != nil {
			s.Emit("error", err.Error())
			s.Close()
			return
		}
		s.Emit("listanswer",data)
	})
	go server.Serve()
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}