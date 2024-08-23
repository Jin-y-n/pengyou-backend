package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pengyou/constant"
	"pengyou/model/common/request"
	"pengyou/model/common/response"
	db "pengyou/storage/database"
	"pengyou/utils/log"
)

func SearchPost(pqi *request.PostQueryInput, c *gin.Context) {

	rows, total, err := db.QueryPost(*pqi)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Logger.Error("rows err", zap.Error(err))
		}
	}(rows)

	if err != nil {
		response.FailWithMessage(constant.SearchFailed, c)
	}

	response.OkWithData(response.PageResult{
		List:     rows,
		Total:    *total,
		PageSize: pqi.PageInfo.PageSize,
		Page:     pqi.PageInfo.Page,
	}, c)

}
