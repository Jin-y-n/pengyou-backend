package service

import (
	"pengyou/constant"
	"pengyou/model"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/utils/check"
	"pengyou/utils/log"
	"pengyou/utils/storage"
	"strings"
	"sync"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upGrade = websocket.Upgrader{}

// establish websocket connection by userId
func EstablishWsConn(c *gin.Context) {
	upGrade.CheckOrigin(c.Request)

	log.Info("upGrade",
		zap.String("subprotocols", strings.Join(upGrade.Subprotocols, ",")),
	)

	userIdStr := "1"
	// c.GetHeader(constant.USER_ID)

	if check.IsBlank(&userIdStr) {
		log.Error("userId is blank")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		log.Error("userId is not a number", zap.Error(err))
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Error("upgrade websocket failed", zap.Error(err))
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

	log.Info("user connect success: " + userIdStr)
}
