package model

import (
	"context"
	"pengyou/constant"

	"pengyou/model/entity"
	rds "pengyou/storage/redis"
	"sync"

	"github.com/gorilla/websocket"
)

// this record the user's node info, include his / her websocket, basic info

type UserNode struct {
	User *entity.User `json:"user"`

	// websocket
	Conn        *websocket.Conn
	Established bool

	// lock
	Lock   *sync.RWMutex
	WsLock *sync.Mutex

	// last heart beat time
	Chatters []string

	// last hadler time
	LastHandlerTime int64
}

var userNodeMap = make(map[string]*UserNode)

func GetUserNode(userId string) *UserNode {
	return userNodeMap[userId]
}

func AddUserNode(userId string, userNode *UserNode) {
	userNodeMap[userId] = userNode
}

func RemoveUserNode(userId string) {
	delete(userNodeMap, userId)
}

func GetUserChatList(userId string) {
	// storage.RedisClient.Get()
}

func AddUserChatList(userId, chatterId string) {
	res := rds.Get(context.Background(), constant.REDIS_USER_CHAT_LIST_PREFIX+userId)

	if res == nil {

	}

}
