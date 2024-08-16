package service

import (
	"pengyou/constant"
	"pengyou/global/config"
	"pengyou/model"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/utils/check"
	"pengyou/utils/log"
	"pengyou/utils/storage"
	"sync"
	"time"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrade = websocket.Upgrader{
	ReadBufferSize:  config.Cfg.Files.ReadBufSize,
	WriteBufferSize: config.Cfg.Files.WriteBufSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// establish websocket connection by userId
func EstablishWsConn(c *gin.Context) {
	userIdStr := c.GetHeader(constant.USER_ID)

	if check.IsBlank(&userIdStr) {
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		response.FailWithMessage(constant.ESTABLISH_WEBSOCKET_CONNECT_FAIL, c)
		return
	}

	user := entity.User{}

	user.ID = uint(userId)

	// add user node record
	storage.AddUserNode(userIdStr,
		&model.UserNode{
			Conn:            ws,
			User:            &user,
			Established:     true,
			Lock:            &sync.RWMutex{},
			WsLock:          &sync.Mutex{},
			LastHandlerTime: time.Time{}.UnixMilli(),
		},
	)

	MsgHandler(storage.GetUserNode(userIdStr))

	log.Logger.Info("user connect success: " + userIdStr)
}
