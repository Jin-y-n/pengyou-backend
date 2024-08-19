package service

import (
	"context"
	"encoding/json"
	"fmt"
	"pengyou/constant"
	"pengyou/model"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/storage"
	rds "pengyou/storage/redis"
	"pengyou/utils/log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upGrade = websocket.Upgrader{}

// establish websocket connection by userId
func EstablishWsConn(c *gin.Context, userId uint) {

	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Error("upgrade websocket failed", zap.Error(err))
		response.FailWithMessage(constant.ESTABLISH_WEBSOCKET_CONNECT_FAIL, c)
		return
	}

	user := entity.User{}

	user.ID = userId

	// add user node record
	storage.AddUserNode(fmt.Sprint(userId),
		&model.UserNode{
			Conn:            ws,
			User:            &user,
			Established:     true,
			Lock:            &sync.RWMutex{},
			WsLock:          &sync.Mutex{},
			LastHandlerTime: time.Time{}.UnixMilli(),
		},
	)

	MsgHandler(storage.GetUserNode(string(userId)))

	log.Info("user connect success: " + string(userId))
}

// cut the connection
func ShutdownWsConn(c *gin.Context, userId uint) {
	// close connection
	user := storage.GetUserNode(string(userId))

	if user == nil {
		log.Error("user not found: " + string(userId))
		response.FailWithMessage(constant.CONNECTED_USER_NOT_FOUND, c)
	}

	if user.Conn == nil {
		log.Error("user not connect: " + string(userId))
	}

	user.Conn.Close()

	// and send the message to all chatters with this user

	log.Info("user disconnect success: " + string(userId))
}

func HeartBeat(c *gin.Context, userId uint) {
	rds.Set(context.Background(), constant.REDIS_USER_HEARTBEAT_PREFIX+string(userId), time.Now().String())
}

func EstablishChatTo(c *gin.Context, from, to uint) {
	userNode := storage.GetUserNode(fmt.Sprint(from))
	if userNode == nil {
		log.Warn("user node not found", zap.String("userId", fmt.Sprint(from)))
	}

}

func CutChat(c *gin.Context, userId, objectId uint) {

	usr := rds.Get(context.Background(), constant.REDIS_USER_CHAT_LIST_PREFIX+string(userId))
	if usr.Val() == "" {
		log.Error("user not found: " + string(userId))
		response.FailWithMessage(constant.CONNECTED_USER_NOT_FOUND, c)
		return
	}
	list := &model.UserChatList{}

	usrBytes, err := usr.Bytes()
	if err != nil {
		log.Error("get bytes error:", zap.Error(err), zap.String("src:", usr.Val()))
		response.FailWithMessage(constant.CHATTER_NOT_FOUND, c)
		return
	}

	err = json.Unmarshal(usrBytes, list)
	if err != nil {
		log.Error("data unmarshal failed", zap.Error(err))
		response.FailWithMessage(constant.SERVER_ERROR, c)
		return
	}

	for i, v := range list.Chatters {
		if v == string(userId) {
			list.Chatters = append(list.Chatters[:i], list.Chatters[i+1:]...)
		}
	}

	chatListBytes, err := json.Marshal(list)
	if err != nil {
		log.Error("data marshal failed", zap.Error(err))
		response.FailWithMessage(constant.SERVER_ERROR, c)
		return
	}

	rds.Set(context.Background(), constant.REDIS_USER_CHAT_LIST_PREFIX+string(userId), string(chatListBytes))

	userNode := storage.GetUserNode(string(objectId))
	if userNode == nil {
		log.Error("user node not found", zap.String("userId", string(objectId)))
		response.FailWithMessage(constant.CHATTER_NOT_FOUND, c)
		return
	}

	userNode.Conn.WriteMessage(websocket.TextMessage, []byte("your connection to "+string(userId)+" has been cut"))
	response.OkWithMessage(constant.CHAT_CUT_SUCCESS, c)
}
