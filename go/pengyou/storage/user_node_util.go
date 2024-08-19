package storage

import (
	"context"
	"encoding/json"
	"pengyou/constant"
	"pengyou/model"
	rds "pengyou/storage/redis"
	"pengyou/utils/log"

	"go.uber.org/zap"
)

// this file records the UserNode of all users's Connect

var userNodeMap = make(map[string]*model.UserNode)

func GetUserNode(userId string) *model.UserNode {

	res := rds.Get(context.Background(), constant.REDIS_USER_NODE_PREFIX+userId)

	if res.Err() != nil {
		log.Error("err", zap.Error(res.Err()))
		return nil
	}

	userNode := &model.UserNode{}

	err := json.Unmarshal([]byte(res.Val()), userNode)

	if err != nil {
		log.Error("unmarshal err: caused by", zap.String("obj:", res.Val()), zap.Error(err))
		return nil
	}

	return userNode
}

func AddUserNode(userId string, userNode *model.UserNode) bool {
	res := rds.SetObj(context.Background(), constant.REDIS_USER_NODE_PREFIX+userId, userNode)

	if res.Err() != nil {
		return false
	}

	return true
}

func RemoveUserNode(userId string) bool {
	res := rds.RedisClient.Del(context.Background(), constant.REDIS_USER_NODE_PREFIX+userId)

	if res.Err() != nil {
		log.Error("remove user node err", zap.String("userId:", userId), zap.Error(res.Err()))
		return false
	}

	return true
}

func GetUserNodeMap() map[string]*model.UserNode {
	keys, err := rds.ScanKeysWithPrefix(constant.REDIS_USER_NODE_PREFIX)

	if err != nil {
		log.Error("scan keys err", zap.Error(err))
	}

	for _, key := range keys {
		res := rds.Get(context.Background(), key)

		if res.Err() != nil {
			log.Error("get user node err", zap.String("key:", key), zap.Error(res.Err()))
		}

		tmp := &model.UserNode{}

		err := json.Unmarshal([]byte(res.Val()), tmp)

		if err != nil {
			log.Error("unmarshal err: caused by", zap.String("obj:", res.Val()), zap.Error(err))
		}

		userNodeMap[key[len(constant.REDIS_USER_NODE_PREFIX):]] = tmp
	}

	return userNodeMap
}
