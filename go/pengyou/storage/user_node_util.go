package storage

import (
	"pengyou/model"
)

// this file records the UserNode of all users's Connect

var userNodeMap = make(map[string]*model.UserNode)

func GetUserNode(userId string) *model.UserNode {
	return userNodeMap[userId]
}

func AddUserNode(userId string, userNode *model.UserNode) bool {
	userNodeMap[userId] = userNode
	return true
}

func RemoveUserNode(userId string) bool {
	delete(userNodeMap, userId)
	return true
}

func GetUserNodeMap() map[string]*model.UserNode {

	return userNodeMap
}
