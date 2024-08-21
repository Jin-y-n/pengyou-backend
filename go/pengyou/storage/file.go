package storage

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"pengyou/constant"
	"pengyou/global/config"
	"pengyou/model/entity"
	db "pengyou/storage/database"
	rds "pengyou/storage/redis"
	"pengyou/utils/log"

	"github.com/gorilla/websocket"
)

// this file implements the function that transfer files

// read file from websocket
func ReadWsFile(ws *websocket.Conn) ([]byte, bool) {
	messageType, buf, err := ws.ReadMessage()

	if err != nil {
		log.Error("read file error" + err.Error())
		return nil, false
	}

	if messageType != websocket.BinaryMessage {
		log.Error("read file error: message type is not binary")
		return nil, false
	}

	return buf, true
}

// write file to websocket
func WriteWsFile(ws *websocket.Conn, buf []byte) bool {
	err := ws.WriteMessage(websocket.BinaryMessage, buf)
	if err != nil {
		log.Error("write file error" + err.Error())
		return false
	}
	return true
}

// save file to file storage
func SaveToFile(w *bufio.Writer, buf []byte) bool {

	nn, err := w.Write(buf)

	if err != nil {
		log.Error("write file error" + err.Error())
		return false
	}

	if nn != len(buf) {
		log.Error("write file error, write bytes not equal to the length of the buffer")
		return false
	}

	return true
}

// read file from file storage (won't close automatically)
func ReadFile(fileName string) (*os.File, bool) {
	file, err := os.OpenFile(config.Cfg.Files.FilePath+fileName, os.O_RDONLY, 0666)

	if err != nil {
		log.Error("open file error" + err.Error())
		return nil, false
	}

	return file, true
}

// create file, each file should have a unique name
// return a file pointer, its default mode is append
func CreateFile(fileName string) (*os.File, bool) {
	file, err := os.Create(config.Cfg.Files.FilePath + fileName)

	if err != nil {
		log.Error("create file error" + err.Error())
		return nil, false
	}

	if err := file.Close(); err != nil {
		log.Error("close file error" + err.Error())
		return nil, false
	}

	file, err = os.OpenFile(config.Cfg.Files.FilePath+fileName, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Error("open file error" + err.Error())
		return nil, false
	}

	return file, true
}

func PersistFile() {

	res, err := rds.GetRedisMemoryUsed()

	if err != nil {
		log.Error("get memory used error" + err.Error())
		return
	}

	if res < config.Cfg.Files.MesToDBThreshold {
		PersistHandledRecord()
	} else {
		PersistAllRecord()
	}
}

func PersistAllRecord() {
	userNodeMap := GetUserNodeMap()

	now := time.Now().UnixMilli()

	for _, userNode := range userNodeMap {
		if userNode.LastHandlerTime+1000 < now {
			res, err := rds.ZRangeByScore(
				context.Background(),
				rds.GenerateName(userNode.User.ID),
				fmt.Sprint(0), fmt.Sprint(now),
			)

			if err != nil {
				log.Error("read message error" + err.Error())
				continue
			}

			for _, message := range res {
				log.Info("read message and saving + " + message)

			}
		}
	}
}

func PersistHandledRecord() {
	userNodeMap := GetUserNodeMap()

	// loop read the nodes

	for _, userNode := range userNodeMap {
		lastTime := userNode.LastHandlerTime

		result, err := rds.ZRangeByScore(context.Background(),
			rds.GenerateName(userNode.User.ID),
			fmt.Sprint(0), fmt.Sprint(lastTime))

		if err != nil {
			log.Error("read message error" + err.Error())
			continue
		}

		// read the messages these are not saved
		for _, message := range result {
			log.Info("read message and saving + " + message)

			splits := strings.Split(message, ",")

			if len(splits) != 3 {
				log.Error("read message error (not full argument) " + message)
				continue
			}

			userId, err := strconv.ParseUint(splits[0], 10, 64)
			if err != nil {
				log.Error("read message error (no sender) " + message)
				continue
			}
			sendTime, err := time.Parse(splits[1], constant.TimeFormatString)
			if err != nil {
				log.Error("read message error (no sending time) " + message)
				continue
			}
			content := splits[2]

			// store the messages to DB and remove the messages from redis
			mesToStore := &entity.MessageSend{
				SenderId:    uint(userId),
				RecipientId: userNode.User.ID,

				Content: content,
				SentAt:  sendTime,

				IsRead: 1,
			}

			db.GormDB.Create(mesToStore)

			mesReceiveToStore := &entity.MessageReceive{
				MessageSendId: mesToStore.ID,
				RecipientId:   mesToStore.RecipientId,

				ReadAt: sendTime,
			}
			db.GormDB.Create(mesReceiveToStore)
		}
	}
}

// create file storage directory
func InitFile(cfg *config.Config) {
	fileDir := cfg.Files.FilePath
	file, err := os.Open(fileDir)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(fileDir, 0755)
			if err != nil {
				panic(err)
			}
		}
	}

	log.Info("file storage directory init success")

	defer file.Close()
}
