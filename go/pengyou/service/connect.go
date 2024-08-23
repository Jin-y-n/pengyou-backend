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

// EstablishWsConn establish websocket connection by userId
func EstablishWsConn(c *gin.Context, userId uint) {

	// upgrade http connection to websocket connection
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Logger.Error("upgrade websocket failed", zap.Error(err))
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

	// storage user connect info to redis
	rds.Set(context.Background(),
		constant.RedisUserEstablishedConnect+strconv.Itoa(int(userId)),
		strconv.Itoa(1))

	log.Logger.Info("user connect success: " + strconv.Itoa(int(userId)))
}

// ShutdownWsConn cut the connection
func ShutdownWsConn(c *gin.Context, userId uint) {
	// close connection
	user := storage.GetUserNode(strconv.Itoa(int(userId)))
	if user != nil {
		if user.Conn == nil {
			log.Logger.Error("user not connect: " + strconv.Itoa(int(userId)))
		}

		wsutil.Close(user.Conn)

		// send the message to all chatters with this user
		if chatter := storage.GetUserNode(strconv.Itoa(int(userId))).Chatters; chatter != nil {

			for _, v := range chatter {
				chatterId, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					log.Logger.Error("parse chatter chatterId error",
						zap.String("user", strconv.Itoa(int(userId))),
						zap.String("chatter", v),
						zap.Error(err))
					continue
				}

				mes := &model.MessageRedis{
					SenderId:    userId,
					RecipientId: uint(chatterId),
					Type:        constant.CutChatMessage,
					SentAt:      time.Now(),
					Content:     constant.RedisCutChatMessageFromPrefix + v,
				}

				err = rds.RedisPublishObj(context.Background(),
					rds.GenerateName(uint(chatterId)),
					mes)
				if err != nil {
					log.Logger.Error("redis publish message failed",
						zap.String("user", strconv.Itoa(int(userId))),
						zap.String("chatter", strconv.Itoa(int(userId))),
						zap.String("msg", constant.RedisDisconnectMessagePrefix+strconv.Itoa(int(userId))),
						zap.Error(err))
				}
			}
		}

		user.Established = false
		log.Logger.Info("user disconnect success: " + strconv.Itoa(int(userId)))
	} else {
		log.Logger.Error("user not found: " + strconv.Itoa(int(userId)))
		response.FailWithMessage(constant.ConnectedUserNotFound, c)
	}

}

// HeartBeat the function that handle the heartbeat
func HeartBeat(c *gin.Context, userId uint) {
	rds.SetWithExpire(context.Background(),
		constant.RedisUserHeartbeatPrefix+strconv.Itoa(int(userId)),
		time.Now().String(),
		constant.HeartBeatTimeout)

	log.Logger.Info("user heartbeat: " + strconv.Itoa(int(userId)))
}

// EstablishChatTo the function that handle the connect link between two users
func EstablishChatTo(c *gin.Context, from, to uint) {
	// get user info
	userNode := storage.GetUserNode(fmt.Sprint(from))
	if userNode != nil {
		if userNode.Established {
			log.Logger.Warn("user node not found", zap.String("userId", fmt.Sprint(from)))
			response.FailWithMessage(constant.RespNotConnect, c)
			return
		}

		msg := model.MessageRedis{
			SenderId:    from,
			RecipientId: to,
			Type:        constant.EstablishChatMessage,
			SentAt:      time.Now(),
			Content:     constant.ChatEstablishSuccessFrom + fmt.Sprint(from),
		}

		// send establish message to target
		err := rds.RedisPublishObj(context.Background(),
			rds.GenerateName(to),
			msg)
		if err != nil {
			log.Logger.Error("establish connect to {" + fmt.Sprint(to) + " }failed")
			response.FailWithMessage(constant.ChatEstablishFailTo+fmt.Sprint(to), c)
		}
		// wait for response from target
		response.OkWithMessage(
			constant.RespWaitForResponse,
			c,
		)
		go func() {
			count := 0
			for {
				time.Sleep(1 * time.Second)

				info := rds.Get(context.Background(),
					constant.RedisUserEstablishedConnectConfirmPrefix+fmt.Sprint(from)+fmt.Sprint(to),
				)

				// receive the confirmation message from redis
				if info != nil {
					if info.Err() != nil {
						log.Logger.Error("read message failed",
							zap.Error(info.Err()))
						break
					}

					if strings.Contains(info.Val(), constant.Yes) {

						userNode.Chatters = append(userNode.Chatters, fmt.Sprint(to))
						response.OkWithMessage(constant.ChatEstablishSuccess, c)

						rds.Del(constant.RedisUserEstablishedConnectConfirmPrefix + fmt.Sprint(from) + fmt.Sprint(to))
						return
					} else if strings.Contains(info.Val(), constant.No) {
						response.OkWithMessage(constant.ChatEstablishFailFrom+fmt.Sprint(to), c)

						rds.Del(constant.RedisUserEstablishedConnectConfirmPrefix + fmt.Sprint(from) + fmt.Sprint(to))
						return
					} else {
						log.Logger.Error("unknown message: ",
							zap.String("msg", info.Val()))
						rds.Del(constant.RedisUserEstablishedConnectConfirmPrefix + fmt.Sprint(from) + fmt.Sprint(to))

						return
					}
				}

				if count > 15 {
					break
				}
				count++
			}
			response.OkWithMessage(constant.RespNoResponse, c)
		}()
	} else {
		log.Logger.Warn("user node not found", zap.String("userId", fmt.Sprint(from)))
		response.FailWithMessage(constant.RespNotConnect, c)
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
						log.Logger.Error(constant.RespCutChatMessageSuccess)
						response.FailWithMessage(constant.ServerError, c)
						return
					}

					response.OkWithMessage(constant.RespCutChatMessageSuccess, c)
					return
				}
			}

		} else {
			log.Logger.Warn("user node not connected", zap.String("userId", fmt.Sprint(userId)))
			response.FailWithMessage(constant.RespNotConnect, c)
		}
	} else {
		log.Logger.Warn("user node not found", zap.String("userId", fmt.Sprint(userId)))
		response.FailWithMessage(constant.RespNotConnect, c)
	}

}

// CheckUserConnect the function that check the user connect
func CheckUserConnect() {

	users := storage.GetUserNodeMap()

	for _, v := range users {
		if !v.Established {
			continue
		}

		if cmd := rds.Get(context.Background(), constant.RedisUserHeartbeatPrefix+strconv.Itoa(int(v.User.ID))); cmd.Err() == nil {
			if t, err := time.Parse(constant.TimeFormatString, cmd.Val()); err == nil {
				if t.Add(constant.HeartBeatTimeout).Before(time.Now()) {
					v.Established = false
					if !wsutil.SendTextMessage(v.Conn, 0, v.User.ID, constant.ConnectCut) {
						log.Logger.Error("send message to " + strconv.Itoa(int(v.User.ID)) + " failed")
					}
				}
			}
		} else if cmd.Err() != nil {
			log.Logger.Error("redis get failed:", zap.Error(cmd.Err()))
			if !wsutil.SendTextMessage(v.Conn, 0, v.User.ID, constant.ConnectCut) {
				log.Logger.Error("send message to " + strconv.Itoa(int(v.User.ID)) + " failed")
			}
		}
	}
}
