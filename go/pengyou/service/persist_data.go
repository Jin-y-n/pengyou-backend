package service

import (
	"pengyou/utils/storage"
)

// this file implement the storage of message from redis to database

func Persist() {
	storage.GetUserNodeMap()
}
