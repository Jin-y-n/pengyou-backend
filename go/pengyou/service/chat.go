package service

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"pengyou/constant"
	"pengyou/global/config"
	"pengyou/model"
	"pengyou/model/common/request"
	"pengyou/storage"
	rds "pengyou/storage/redis"
	chatutil "pengyou/utils/chat"
	"pengyou/utils/log"
	strutil "pengyou/utils/string"
	wsutil "pengyou/utils/ws"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var mesDispatchRule = make(map[string]func(message string))

// this file implements the chat function
func MsgHandler(userNode *model.UserNode) {

	// check the connect
	ws := userNode.Conn
	if config.Cfg == nil || config.Cfg.App.PublishKey == "" {
		ws.WriteMessage(websocket.TextMessage, []byte(constant.SERVER_ERROR))
		log.Error("PublishKey is not configured")
		return
	}

	go MsgPublish(ws, userNode)
	go MsgSubscribe(ws, userNode)
}

// MsgPublish publish the message to redis
func MsgPublish(ws *websocket.Conn, userNode *model.UserNode) {

	// handler message
	for userNode.Established {
		func() {

			// read message
			message := &request.MessageIn{}
			err := ws.ReadJSON(message)

			if err != nil {
				log.Warn("read ws message error:" + err.Error())
				if strings.Contains(err.Error(), "websocket: close") {
					userNode.Established = false
					return
				}
			}
			// check the send time of the message is valid or not
			if math.Abs(float64(message.CreateAt.UnixMilli()-time.Now().UnixMilli())) > 1000 {
				log.Warn("message time error")
				ws.WriteMessage(websocket.TextMessage, []byte("message time error, please check your network and try again"))
				return
			}
			log.Info("read ws message:" + string(message.Content))

			switch message.RequestType {
			case constant.MESSAGE_TYPE_TEXT:
				publishText(message)
			case constant.MESSAGE_TYPE_FILE_REQUEST:
				success := uploadFile(ws, message.Content)
				if !success {
					ws.WriteMessage(websocket.TextMessage,
						[]byte("upload file error, please try again"))
					return
				}

				ws.WriteMessage(websocket.TextMessage,
					[]byte("upload file success"))
				publishText(message)
			}
		}()
	}

	defer func() {
		ws.Close()
		log.Info("close websocket")
	}()
}

func MsgSubscribe(ws *websocket.Conn, userNode *model.UserNode) {
	for userNode.Established {
		func() {

			// get unhandled messages
			now := time.Now().UnixMilli()
			result, err := rds.ZRangeByScore(
				context.Background(),
				rds.GenerateName(userNode.User.ID),
				fmt.Sprint(float64(userNode.LastHandlerTime)),
				fmt.Sprint(float64(now)))

			if err != nil {
				log.Warn("subscribing message error:" + err.Error())

				if strings.Contains(err.Error(), "websocket: close") {
					userNode.Established = false
					return
				}
			} else {
				// send unhandled messages
				for _, message := range result {

					// TODO: more message types
					// MessageDispatcher(message, map[string]func(message string))
					if strings.HasPrefix(message, constant.REDIS_DISCONNECT_MESSAGE_PREFIX) {

						from := strings.TrimPrefix(message, constant.REDIS_DISCONNECT_MESSAGE_PREFIX)
						userNode.Chatters = strutil.RemoveElementByValue(userNode.Chatters, from)
						wsutil.SendTextMessage(ws, constant.RESP_DISCONNECT_MESSAGE_PREFIX+from)

					} else if strings.HasPrefix(message, constant.REDIS_ESTABLISH_CHAT_MESSAGE_FROM_PREFIX) {

						go func() {
							from := strings.TrimPrefix(message, constant.REDIS_ESTABLISH_CHAT_MESSAGE_FROM_PREFIX)

							wsutil.SendTextMessage(ws, constant.RESP_ESTABLISH_CHAT_MESSAGE_FROM_PREFIX+from)

							chatutil.AddEstablishRequestNode(from, strconv.Itoa(int(userNode.User.ID)))
							count := 1
							for {
								time.Sleep(1 * time.Second)

								if count < 6 && chatutil.GetEstablishRequestNode(from, strconv.Itoa(int(userNode.User.ID))) {
									chatutil.RemoveEstablishRequestNode(from, strconv.Itoa(int(userNode.User.ID)))
									userNode.Chatters = append(userNode.Chatters, from)
									wsutil.SendTextMessage(ws, constant.CHAT_ESTABLISH_SUCCESS_FROM+from)

									return
								}

								if count >= 6 {
									wsutil.SendTextMessage(ws, constant.CHAT_ESTABLISH_FAIL_FROM+from)
									return
								}
								count++
							}
						}()
					} else if strings.HasPrefix(message, constant.REDIS_CUT_CHAT_MESSAGE_FROM_PREFIX) {
						from := strings.TrimPrefix(message, constant.REDIS_CUT_CHAT_MESSAGE_FROM_PREFIX)

						userNode.Chatters = strutil.RemoveElementByValue(userNode.Chatters, from)
						wsutil.SendTextMessage(ws, constant.RESP_CHAT_CUTTED_FROM+from)
					} else if (strings.HasPrefix(message, constant.REDIS_USER_DISCONNECT)) {

						from := strings.TrimPrefix(message, constant.REDIS_USER_DISCONNECT)

						userNode.Chatters = strutil.RemoveElementByValue(userNode.Chatters, from)
						wsutil.SendTextMessage(ws, constant.RESP_CHATTRT_DISCONNECTED+from)
					} else {
						wsutil.SendTextMessage(ws, message)
					}

					userNode.LastHandlerTime = now
				}
			}

		}()
	}

	defer func() {
		ws.Close()
		log.Info("close websocket")
	}()
}

func publishText(message *request.MessageIn) {
	messageRedis := model.MessageRedis{
		Content:     message.Content,
		RecipientId: message.RecipientId,
		SenderId:    message.SenderId,
		SentAt:      time.Now(),
		Type:        constant.MESSAGE_TYPE_TEXT,
	}

	// send message
	rds.RedisClient.ZAdd(context.Background(),
		rds.GenerateName(message.RecipientId),
		redis.Z{
			Score:  float64(time.Now().UnixMilli()),
			Member: messageRedis})

	log.Info("publish message:" + string(message.Content))
}

func uploadFile(ws *websocket.Conn, fileName string) bool {
	log.Info("uploading file (" + fileName + ") ...")

	file, success := storage.CreateFile(fileName)
	if !success {
		log.Warn("create file error: " + fileName)
	}
	// w := bufio.NewWriter(file)

	defer file.Close()

	// loop read
	loop := true
	for loop {
		buf, success := storage.ReadWsFile(ws)

		if len(buf) != int(config.Cfg.Files.ReadBufSize) {
			loop = false
		}

		if !success {
			log.Warn("read file error: " + fileName)
			return false
		}

		// if !storage.SaveToFile(w, buf) {
		// 	log.Warn("save file error: " + fileName)
		// 	return false
		// }
	}

	return true
}

func downloadFile(ws *websocket.Conn, fileName string) {
	if file, success := storage.ReadFile(fileName); success {
		defer file.Close()

		r := bufio.NewReader(file)
		buf := make([]byte, config.Cfg.Files.ReadBufSize)

		n, err := r.Read(buf)
		for n != 0 {
			if err != nil {
				log.Warn("read file error: " + fileName)
			}

			// if !storage.WriteWsFile(ws, buf) {
			// 	log.Warn("write file error: " + fileName)
			// }
		}
	} else {
		wsutil.SendTextMessage(ws, "file not found")
	}
}

// this function is designed to dispatch the messages subscribed from redis
func MessageDispatcher(message string, rule map[string]func(message string)) {
	for k, v := range rule {
		if strings.HasPrefix(message, k) {
			v(message)
		}
	}
}
