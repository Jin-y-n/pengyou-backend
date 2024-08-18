package storage

import "pengyou/model"

// this file records the UserNode of all users's Connect

var userNodeMap = make(map[string]*model.UserNode)

func GetUserNode(userId string) *model.UserNode {
	return userNodeMap[userId]
}

func AddUserNode(userId string, userNode *model.UserNode) {
	userNodeMap[userId] = userNode
}

func RemoveUserNode(userId string) {
	delete(userNodeMap, userId)
}

func GetUserNodeMap() map[string]*model.UserNode {
	return userNodeMap
}
