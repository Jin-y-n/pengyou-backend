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
	"pengyou/utils/log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

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

// publish the message to redis
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
			// TODO: uncommit this part to confirm the message time --------------------------------------------
			// check the send time of the message is valid or not
			// if math.Abs(float64(message.CreateAt.UnixMilli()-time.Now().UnixMilli())) > 1000 {
			// 	log.Warn("message time error")
			// 	ws.WriteMessage(websocket.TextMessage, []byte("message time error, please check your network and try again"))
			// 	return
			// }
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

			math.Abs(1)

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
					log.Info("sending message:" + message)
					ws.WriteMessage(websocket.TextMessage, []byte(message))

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
	w := bufio.NewWriter(file)

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

		if !storage.SaveToFile(w, buf) {
			log.Warn("save file error: " + fileName)
			return false
		}
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

			if !storage.WriteWsFile(ws, buf) {
				log.Warn("write file error: " + fileName)
			}
		}
	} else {
		ws.WriteMessage(websocket.TextMessage, []byte("file not found"))
	}
}
