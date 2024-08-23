package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pengyou/constant"
	"pengyou/model/common/request"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/service"
	"pengyou/utils/common"
	"strconv"
)

// AddPost @Summary 添加帖子
// @Description 添加新的帖子
// @Tags 帖子
// @Accept json
// @Produce json
// @Param post body request.PostCreateInput true "帖子信息"
// @Success 200 {object} response.Response "成功响应"
// @Failure 400 {object} response.Response "无效参数"
// @Router /posts [post]
func AddPost(c *gin.Context) {
	post := &request.PostCreateInput{}

	err := c.ShouldBindJSON(post)

	if err != nil {
		response.FailWithMessage(constant.InvalidParams, c)
		return
	}

	author := post.Author
	authorId, _ := strconv.Atoi(author)

	if !common.CheckUserIdDefault(uint(authorId)) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	postEntity := &entity.Post{
		Author:  uint(authorId),
		Content: post.Content,
		Title:   post.Title,
		Model: gorm.Model{
			CreatedAt: post.CreatedAt,
			ID:        uint(common.NextSnowflakeID()),
		},
	}

	service.AddPost(postEntity, c)
}

// UpdatePost @Summary 更新帖子
// @Description 更新帖子的信息
// @Tags 帖子
// @Accept json
// @Produce json
// @Param post body request.PostUpdateInput true "帖子信息"
// @Success 200 {object} response.Response "成功响应"
// @Failure 400 {object} response.Response "无效参数"
// @Router /posts [put]
func UpdatePost(c *gin.Context) {
	post := &request.PostUpdateInput{}

	err := c.ShouldBindJSON(post)

	if err != nil {
		response.FailWithMessage(constant.InvalidParams, c)
		return
	}

	author := post.Author
	authorId, _ := strconv.Atoi(author)

	if !common.CheckUserIdDefault(uint(authorId)) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	postEntity := &entity.Post{
		Author:  uint(authorId),
		Content: post.Content,
		Title:   post.Title,
		Model: gorm.Model{
			UpdatedAt: post.UpdatedAt,
			ID:        uint(common.NextSnowflakeID()),
		},
	}

	service.UpdatePost(postEntity, c)
}

// DeletePost @Summary 删除帖子
// @Description 删除指定 ID 的帖子
// @Tags 帖子
// @Accept json
// @Produce json
// @Param id body request.GetById true "帖子 ID"
// @Success 200 {object} response.Response "成功响应"
// @Failure 400 {object} response.Response "无效参数"
// @Router /posts [delete]
func DeletePost(c *gin.Context) {
	post := &request.GetById{}

	err := c.ShouldBindJSON(post)

	if err != nil {
		response.FailWithMessage(constant.InvalidParams, c)
		return
	}

	service.DeletePost(post.ID, c)
}
