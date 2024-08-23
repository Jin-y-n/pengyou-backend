package storage

import "pengyou/model"

var rdsMsgReadList = make(map[string]model.MessageRedis)

func GetMsgReadList() map[string]model.MessageRedis {
	return rdsMsgReadList
}

func SetMsgReadList(key string, value model.MessageRedis) {
	rdsMsgReadList[key] = value
}

func DelMsgReadList(key string) {
	delete(rdsMsgReadList, key)
}

func GetMsgReadListByKey(key string) model.MessageRedis {
	return rdsMsgReadList[key]
}

func AddMsgReadList(key string, value model.MessageRedis) {
	rdsMsgReadList[key] = value
}
