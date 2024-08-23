package service

import (
	"pengyou/constant"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/utils/log"
	"strconv"

	es "pengyou/storage/elasticsearch"

	"github.com/gin-gonic/gin"
)

func AddPost(post *entity.Post, c *gin.Context) {

	log.Logger.Info("post: " + post.Content)
	log.Logger.Info("title: " + post.Title)
	log.Logger.Info("author: " + strconv.Itoa(int(post.Author)))

	err := es.IndexPostAdd(post)
	if err != nil {
		response.FailWithMessage(constant.AddedFailed, c)
		return
	}

	response.OkWithMessage(constant.AddedSuccess, c)
}

func UpdatePost(post *entity.Post, c *gin.Context) {

	err := es.IndexPostUpdate(post)
	if err != nil {
		response.FailWithMessage(constant.UpdatedFailed, c)
		return
	}

	response.OkWithMessage(constant.UpdatedSuccess, c)
}

func DeletePost(post int, c *gin.Context) {

	err := es.IndexPostDelete(post)
	if err != nil {
		response.FailWithMessage(constant.DeletedFailed, c)
		return
	}

	response.OkWithMessage(constant.DeletedSuccess, c)
}
