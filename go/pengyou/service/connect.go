package service

import (
	"context"
	"fmt"
	"pengyou/constant"
	"pengyou/model"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/storage"
	rds "pengyou/storage/redis"
	"pengyou/utils/log"
	strutil "pengyou/utils/string"
	wsutil "pengyou/utils/ws"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upGrade = websocket.Upgrader{}

// establish websocket connection by userId
func EstablishWsConn(c *gin.Context, userId uint) {

	// upgrade http connection to websocket connection
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Error("upgrade websocket failed", zap.Error(err))
		response.FailWithMessage(constant.ESTABLISH_WEBSOCKET_CONNECT_FAIL, c)
		return
	}

	// storage info to user node
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

	// begin handlering message
	MsgHandler(storage.GetUserNode(string(userId)))

	log.Info("user connect success: " + string(userId))
}

// cut the connection
func ShutdownWsConn(c *gin.Context, userId uint) {
	// close connection
	user := storage.GetUserNode(string(userId))

	if user != nil {
		if user.Conn == nil {
			log.Error("user not connect: " + string(userId))
		}

		user.Conn.Close()

		// and send the message to all chatters with this user

		if chatter := storage.GetUserNode(string(userId)).Chatters; chatter != nil {

			for _, v := range chatter {
				chatterId, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					log.Error("parse chatter chatterId error: " + v)
					continue
				}

				rds.RedisPublish(context.Background(), rds.GenerateName(uint(chatterId)), constant.REDIS_DISCONNECT_MESSAGE_PREFIX+string(userId))
			}
		}

		log.Info("user disconnect success: " + string(userId))
	} else {
		log.Error("user not found: " + string(userId))
		response.FailWithMessage(constant.CONNECTED_USER_NOT_FOUND, c)
	}

}

func HeartBeat(c *gin.Context, userId uint) {
	rds.SetWithExpire(context.Background(), constant.REDIS_USER_HEARTBEAT_PREFIX+string(userId), time.Now().String(), constant.HEART_BEAT_TIMEOUT)

	log.Info("user heartbeat: " + string(userId))
}

func EstablishChatTo(c *gin.Context, from, to uint) {
	userNode := storage.GetUserNode(fmt.Sprint(from))
	if userNode == nil || !userNode.Established {
		log.Warn("user node not found", zap.String("userId", fmt.Sprint(from)))
	}

	rds.RedisPublish(context.Background(),
		constant.REDIS_ESTABLISH_CHAT_MESSAGE_FROM_PREFIX+fmt.Sprint(from),
		fmt.Sprint(to))
	userNode.Chatters = append(userNode.Chatters, fmt.Sprint(to))

	response.OkWithMessage(constant.CHAT_ESTABLISH_SUCCESS, c)
}

// the function that handle the connect link between two users
func CutChat(c *gin.Context, userId, chatterId uint) {

	userNode := storage.GetUserNode(strconv.Itoa(int(userId)))
	if userNode == nil || !userNode.Established {
		log.Warn("user node not found", zap.String("userId", fmt.Sprint(userId)))
	}

	chatter := strconv.Itoa(int(chatterId))

	for _, v := range userNode.Chatters {
		if v == chatter {
			userNode.Chatters = strutil.RemoveElementByValue(userNode.Chatters, v)

			rds.RedisPublish(context.Background(), rds.GenerateName(chatterId), constant.REDIS_CUT_CHAT_MESSAGE_FROM_PREFIX+chatter)

			wsutil.SendTextMessage(userNode.Conn, constant.CUT_CHAT_MESSAGE_RESPONSE_SUCCESS)
			return
		}
	}
}

func CheckUserConnect() {
	keysWithPrefix, err := rds.ScanKeysWithPrefix(constant.REDIS_USER_HEARTBEAT_PREFIX)
	if err != nil {
		return
	}

	for _, v := range keysWithPrefix {
		if cmd := rds.Get(context.Background(), v); cmd.Err() == nil {
			if t, err := time.Parse(constant.TIME_FORMAT_STRING, cmd.Val()); err == nil {
				if t.Add(constant.HEART_BEAT_TIMEOUT).Before(time.Now()) {
					userNode := storage.GetUserNode(strings.Trim(v, constant.REDIS_USER_HEARTBEAT_PREFIX))
					if userNode != nil {
						wsutil.SendTextMessage(userNode.Conn, constant.CONNECT_CUTTED)

						chatter := userNode.Chatters
						for _, v := range chatter {
							chatterId, err := strconv.ParseUint(v, 10, 64)
							if err != nil {
								log.Error("convert string to int failed : ", zap.String("arg", v))
							} else {
								rds.RedisPublish(context.Background(), rds.GenerateName(uint(chatterId)), constant.REDIS_USER_DISCONNECT+v)
							}

						}
						rds.Del(context.Background(), v)

						userNode.Established = false
					}

				}
			}
		}
	}
}
