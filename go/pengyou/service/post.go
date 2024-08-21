package service

// import (
// 	"pengyou/constant"
// 	"pengyou/model/common/response"
// 	"pengyou/utils/log"

// 	"github.com/gin-gonic/gin"
// )

// func PostUpload(c *gin.Context) {
// 	post, success := c.GetPostForm(constant.POSTED_CONTENT)

// 	if !success {
// 		response.FailWithMessage(constant.CANNOT_FOUND_CONTENT, c)
// 		return
// 	}

// 	title, success := c.GetPostForm(constant.POSTED_TITLE)

// 	if !success {
// 		response.FailWithMessage(constant.CANNOT_FOUND_TITLE, c)
// 		return
// 	}

// 	user, success := c.GetPostForm(constant.POSTED_USER)
// 	if !success {
// 		response.FailWithMessage(constant.CANNOT_FOUND_USER, c)
// 		return
// 	}

// 	log.Info("post: " + post)
// 	log.Info("title: " + title)
// 	log.Info("user: " + user)

// 	response.OkWithMessage("upload success", c)
// }
