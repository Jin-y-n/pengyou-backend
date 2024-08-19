package model

import (
	"pengyou/model/entity"
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

func (un *UserNode) GetUserChatList() []string {
	return un.Chatters
}

func (un *UserNode) AddUserChatList(chatterId string) {
	un.Chatters = append(un.Chatters, chatterId)
}
