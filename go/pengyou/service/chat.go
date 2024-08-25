package service

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"pengyou/constant"
	"pengyou/global/config"
	"pengyou/model"
	"pengyou/model/common/request"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/storage"
	db "pengyou/storage/database"
	rds "pengyou/storage/redis"
	chatutil "pengyou/utils/chat"
	"pengyou/utils/common"
	fileutil "pengyou/utils/file"
	"pengyou/utils/log"
	strutil "pengyou/utils/string"
	wsutil "pengyou/utils/ws"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var mesDispatchRule = make(map[string]func(message string))

// MsgHandler this file implements the chat function
func MsgHandler(userNode *model.UserNode) {

	// check connect
	ws := userNode.Conn
	if config.Cfg == nil || config.Cfg.App.PublishKey == "" {
		wsutil.SendTextMessage(ws, constant.SystemId, userNode.User.ID, constant.ServerError)
		log.Logger.Error("PublishKey is not configured")
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

			message := MsgReceive(ws, userNode)

			if message == nil {
				log.Logger.Error("read message from websocket failed")
				return
			}
			log.Logger.Info("read ws message:" + string(message.Content))
			message.CreateAt = time.Now()

			MsgPublisher(ws, 0, message)
		}()
	}

	defer func() {
		wsutil.Close(ws)
		log.Logger.Info("close websocket")
	}()
}

// MsgReceive receive the message from websocket
func MsgReceive(ws *websocket.Conn, userNode *model.UserNode) *request.MessageTransfer {
	// read message
	message := &request.MessageTransfer{}
	err := ws.ReadJSON(message)

	if err != nil {
		log.Logger.Warn("read ws message error:" + err.Error())
		if strings.Contains(err.Error(), "websocket: close") {
			userNode.Established = false
			return nil
		}
	}
	// check the send time of the message is valid or not
	//if math.Abs(float64(message.CreateAt.UnixMilli()-time.Now().UnixMilli())) > 1000 {
	//	log.Logger.Warn("message time error")
	//	wsutil.SendTextMessage(ws, constant.SystemId, userNode.User.ID, "message time error, please check your network and try again")
	//	return nil
	//}

	return message
}

// MsgPublisher publish the message to redis
func MsgPublisher(ws *websocket.Conn, user uint, message *request.MessageTransfer) {

	msgRds := model.MessageRedis{
		ID:          uint(common.NextSnowflakeID()),
		SenderId:    message.SenderId,
		RecipientId: message.RecipientId,
		Type:        message.Type,
		SentAt:      message.CreateAt,
		ReceiveAt:   time.Time{},
		Content:     message.Content,
	}

	switch message.Type {
	case constant.MessageTypeText:
		publishText(msgRds)
	case constant.MessageTypeFileRequest:
		success := uploadFile(ws, message.Content)
		if !success {
			wsutil.SendTextMessage(
				ws,
				constant.SystemId,
				user,
				"upload file error, please try again",
			)
			return
		}

		wsutil.SendTextMessage(
			ws,
			constant.SystemId,
			user,
			"upload file success",
		)
		publishText(msgRds)
	}

	// return to the frontend that msg has sent
	wsutil.SendTextMessage(
		ws,
		constant.SystemId,
		user,
		message.Content,
	)
}

// MsgSubscribe subscribe the message from redis
func MsgSubscribe(ws *websocket.Conn, userNode *model.UserNode) {
	for userNode.Established {
		func() {

			// get unhandled messages
			// now := time.Now().UnixMilli()

			err := rds.RedisSubscribe(
				context.Background(),
				rds.GenerateName(userNode.User.ID),
				func(message string) {

					MessageDispatcher(
						ws,
						message,
						userNode)
				})
			if err != nil {
				return
			}
		}()
	}
	defer func() {
		wsutil.Close(ws)
		log.Logger.Info("close websocket")
	}()
}

// publishText publish the message to redis
func publishText(msg model.MessageRedis) {

	// send message
	err := rds.RedisPublishObj(
		context.Background(),
		rds.GenerateName(msg.RecipientId),
		msg)
	if err != nil {
		return
	}

	log.Logger.Info("publish message:" + string(msg.Content))
}

// uploadFile upload the file to local file system
func uploadFile(ws *websocket.Conn, fileName string) bool {
	log.Logger.Info("uploading file (" + fileName + ") ...")

	file, success := storage.CreateFile(fileName)
	if !success {
		log.Logger.Warn("create file error: " + fileName)
	}
	// w := bufio.NewWriter(file)

	defer fileutil.Close(file)

	// loop read
	loop := true
	for loop {
		buf, success := storage.ReadWsFile(ws)

		if len(buf) != int(config.Cfg.Files.ReadBufSize) {
			loop = false
		}

		if !success {
			log.Logger.Warn("read file error: " + fileName)
			return false
		}

		// if !storage.SaveToFile(w, buf) {
		// 	log.Logger.Warn("save file error: " + fileName)
		// 	return false
		// }
	}

	return true
}

// downloadFile download the file from local file system
func downloadFile(ws *websocket.Conn, user uint, fileName string) {
	if file, success := storage.ReadFile(fileName); success {
		defer fileutil.Close(file)

		r := bufio.NewReader(file)
		buf := make([]byte, config.Cfg.Files.ReadBufSize)

		n, err := r.Read(buf)
		for n != 0 {
			if err != nil {
				log.Logger.Warn("read file error: " + fileName)
			}

			// if !storage.WriteWsFile(ws, buf) {
			// 	log.Logger.Warn("write file error: " + fileName)
			// }
		}
	} else {
		wsutil.SendTextMessage(ws, constant.SystemId, user, "file not found")
	}
}

// MessageDispatcher this function is designed to dispatch the messages subscribed from redis
func MessageDispatcher(ws *websocket.Conn, message string, userNode *model.UserNode) bool {

	// unmarshal message to object
	msg := &model.MessageRedis{}
	err := json.Unmarshal([]byte(message), msg)
	if err != nil {
		log.Logger.Error("unmarshal message error: " + err.Error())
		return false
	}

	// store the message to database
	go func() {
		err := db.StoreRdsMessage(userNode, msg)
		if err != nil {
			wsutil.SendTextMessage(
				userNode.Conn,
				constant.SystemId,
				userNode.User.ID,
				constant.MessageStorageFailed)
			return
		}
	}()

	// receive message from user1 who cut chat
	if msg.Type == constant.MessageTypeCutChat {

		// get user1 id and remove user1 from user2's chatting list
		from := msg.SenderId
		userNode.Chatters = strutil.RemoveElementByValue(
			userNode.Chatters,
			strconv.Itoa(int(from)),
		)

		// send message to user2 to cut chat
		wsutil.SendTextMessage(
			ws,
			constant.SystemId,
			userNode.User.ID,
			constant.RespDisconnectMessagePrefix+strconv.Itoa(int(from)),
		)
		// receive message from user1 who wants to establish chat
	} else if msg.Type == constant.MessageTypeEstablishChat {
		// run a routine to establish chat
		go func() {

			// send to user2 the request from user1
			from := msg.SenderId
			wsutil.SendTextMessage(
				ws,
				constant.SystemId,
				userNode.User.ID,
				constant.RespEstablishChatMessageFromPrefix+strconv.Itoa(int(from)),
			)

			// add request in requestNodes
			chatutil.AddEstablishRequestNode(
				strconv.Itoa(int(from)),
				strconv.Itoa(int(userNode.User.ID)),
			)

			// if more than 15 seconds, cancel and send failed message to sender
			count := 1
			for {
				time.Sleep(1 * time.Second)

				// if request is still in requestNodes and is set to true, send success message to sender
				if count < 16 && chatutil.GetEstablishRequestNode(
					strconv.Itoa(int(from)),
					strconv.Itoa(int(userNode.User.ID)),
				) {
					// add chatter to userNode and send the message to user
					userNode.Chatters = append(userNode.Chatters, strconv.Itoa(int(from)))
					wsutil.SendTextMessage(
						ws,
						constant.SystemId,
						userNode.User.ID,
						constant.ChatEstablishSuccessFrom+strconv.Itoa(int(from)),
					)

					// confirmed -> remove the requestNode
					chatutil.RemoveEstablishRequestNode(
						strconv.Itoa(int(from)),
						strconv.Itoa(int(userNode.User.ID)),
					)
					return
				}

				// time out send failed message to two users
				if count > 15 {
					wsutil.SendTextMessage(
						ws,
						constant.SystemId,
						userNode.User.ID,
						constant.ChatEstablishFailFrom+strconv.Itoa(int(from)),
					)

					wsutil.SendTextMessage(
						ws,
						constant.SystemId,
						from,
						constant.ChatEstablishFailTo+strconv.Itoa(int(from)),
					)

					return
				}
				count++
			}
		}()
		// receive the cut chat message
	} else if msg.Type == constant.MessageTypeDisconnect {

		from := msg.SenderId

		userNode.Chatters = strutil.RemoveElementByValue(userNode.Chatters, strconv.Itoa(int(from)))
		wsutil.SendTextMessage(
			ws,
			constant.SystemId,
			userNode.User.ID,
			constant.RespChatterDisconnected+strconv.Itoa(int(from)),
		)
	} else {
		wsutil.SendJsonMsg(ws, &message)
	}

	return true
}

func LeaveMsg(msg *entity.MessageSend, c *gin.Context) {
	msg.ID = uint(common.NextSnowflakeID())

	err := db.StoreMsgSend(msg)

	if err != nil {
		response.FailWithMessage(constant.ServerError, c)
		return
	}

	response.OkWithMessage(constant.SendMsgSuccess, c)
}

func GetMsgUnRead(id uint, c *gin.Context) {

	msg, err := db.QueryUnReadMsgById(id)

	if err != nil {
		response.FailWithMessage(constant.ServerError, c)
		return
	}

	response.OkWithData(msg, c)
}
