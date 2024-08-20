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
		response.FailWithMessage(constant.EstablishWebsocketConnectFail, c)
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

	// begin handling message
	MsgHandler(storage.GetUserNode(strconv.Itoa(int(userId))))

	log.Info("user connect success: " + strconv.Itoa(int(userId)))
}

// ShutdownWsConn cut the connection
func ShutdownWsConn(c *gin.Context, userId uint) {
	// close connection
	user := storage.GetUserNode(strconv.Itoa(int(userId)))

	if user != nil {
		if user.Conn == nil {
			log.Error("user not connect: " + strconv.Itoa(int(userId)))
		}

		wsutil.Close(user.Conn)

		// and send the message to all chatters with this user

		if chatter := storage.GetUserNode(strconv.Itoa(int(userId))).Chatters; chatter != nil {

			for _, v := range chatter {
				chatterId, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					log.Error("parse chatter chatterId error: " + v)
					continue
				}

				err = rds.RedisPublish(context.Background(), rds.GenerateName(uint(chatterId)), constant.RedisDisconnectMessagePrefix+strconv.Itoa(int(userId)))
				if err != nil {
					log.Error("redis publish message failed : ",
						zap.String("mess:", constant.RedisDisconnectMessagePrefix+strconv.Itoa(int(userId))),
						zap.Error(err))
				}
			}
		}

		log.Info("user disconnect success: " + strconv.Itoa(int(userId)))
	} else {
		log.Error("user not found: " + strconv.Itoa(int(userId)))
		response.FailWithMessage(constant.ConnectedUserNotFound, c)
	}

}

func HeartBeat(c *gin.Context, userId uint) {
	rds.SetWithExpire(context.Background(), constant.RedisUserHeartbeatPrefix+strconv.Itoa(int(userId)), time.Now().String(), constant.HeartBeatTimeout)

	log.Info("user heartbeat: " + strconv.Itoa(int(userId)))
}

func EstablishChatTo(c *gin.Context, from, to uint) {
	userNode := storage.GetUserNode(fmt.Sprint(from))
	if userNode != nil {
		if userNode.Established {
			log.Warn("user node not found", zap.String("userId", fmt.Sprint(from)))
			response.FailWithMessage(constant.RespNotConnect, c)
			return
		}

		err := rds.RedisPublish(context.Background(),
			constant.RedisEstablishChatMessageFromPrefix+fmt.Sprint(from),
			fmt.Sprint(to))
		if err != nil {
			log.Error("establish connect to {" + fmt.Sprint(to) + " }failed")
			response.FailWithMessage(constant.ChatEstablishFailTo+fmt.Sprint(to), c)
		}
		userNode.Chatters = append(userNode.Chatters, fmt.Sprint(to))

		response.OkWithMessage(constant.ChatEstablishSuccess, c)
	} else {
		log.Warn("user node not found", zap.String("userId", fmt.Sprint(from)))

	}
}

// CutChat the function that handle the connect link between two users
func CutChat(c *gin.Context, userId, chatterId uint) {

	userNode := storage.GetUserNode(strconv.Itoa(int(userId)))
	if userNode != nil {
		if userNode.Established {
			chatter := strconv.Itoa(int(chatterId))

			for _, v := range userNode.Chatters {
				if v == chatter {
					userNode.Chatters = strutil.RemoveElementByValue(userNode.Chatters, v)

					err := rds.RedisPublish(context.Background(), rds.GenerateName(chatterId), constant.RedisCutChatMessageFromPrefix+chatter)
					if err != nil {
						log.Error(constant.CutChatMessageResponseSuccess)
						return
					}

					wsutil.SendTextMessage(userNode.Conn, constant.CutChatMessageResponseSuccess)
					return
				}
			}

		} else {
			log.Warn("user node not connected", zap.String("userId", fmt.Sprint(userId)))

		}
	} else {
		log.Warn("user node not found", zap.String("userId", fmt.Sprint(userId)))
	}

}

func CheckUserConnect() {
	keysWithPrefix, err := rds.ScanKeysWithPrefix(constant.RedisUserHeartbeatPrefix)
	if err != nil {
		return
	}

	for _, v := range keysWithPrefix {
		if cmd := rds.Get(context.Background(), v); cmd.Err() == nil {
			if t, err := time.Parse(constant.TimeFormatString, cmd.Val()); err == nil {
				if t.Add(constant.HeartBeatTimeout).Before(time.Now()) {
					userNode := storage.GetUserNode(strings.Trim(v, constant.RedisUserHeartbeatPrefix))
					if userNode != nil {
						wsutil.SendTextMessage(userNode.Conn, constant.ConnectCut)

						chatter := userNode.Chatters
						for _, v := range chatter {
							chatterId, err := strconv.ParseUint(v, 10, 64)
							if err != nil {
								log.Error("convert string to int failed : ", zap.String("arg", v))
							} else {
								err := rds.RedisPublish(context.Background(), rds.GenerateName(uint(chatterId)), constant.RedisUserDisconnect+v)
								if err != nil {
									log.Error("publish message failed: ", zap.String("mes: ", constant.RedisUserDisconnect+v))
									return
								}
							}

						}
						rds.Del(v)
						userNode.Established = false
						storage.RemoveUserNode(strings.Trim(v, constant.RedisUserHeartbeatPrefix))
					}

				}
			}
		}
	}
}
