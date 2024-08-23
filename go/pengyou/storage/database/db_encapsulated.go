package db

import (
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pengyou/constant"
	"pengyou/model"
	"pengyou/model/common/request"
	"pengyou/model/entity"
	"pengyou/utils/chat"
	"pengyou/utils/log"
	"time"
)

// StoreRdsMessage StoreMsgToDB this implement the storage of message from redis to database
func StoreRdsMessage(userNode *model.UserNode, message *model.MessageRedis) error {
	mr := entity.MessageReceive{
		ID:              message.ID,
		MessageSenderId: message.SenderId,
		RecipientId:     message.RecipientId,
		ReadAt:          time.Now(),
		DeleteAt:        gorm.DeletedAt{},
		Type:            chat.ConvertFrontendMsgType(message.Type),
	}

	res := GormDB.Save(mr)
	return res.Error
}

func StoreMsgSend(message *entity.MessageSend) error {
	res := GormDB.Save(message)
	return res.Error
}

func StoreMsgRec(id uint, confirmTime time.Time) error {

	msg := &entity.MessageReceive{
		ID:     id,
		ReadAt: confirmTime,
	}

	res := GormDB.Save(msg)

	return res.Error
}

func StoreMsgConfirm(msgId uint) error {
	msg := &entity.MessageSend{
		ID: msgId,
	}

	res := GormDB.Save(msg).Update("is_read = ?", constant.TRUE)

	return res.Error
}

func QueryMsgRecById(id uint) (*entity.MessageReceive, error) {
	msgRec := &entity.MessageReceive{
		ID: id,
	}

	res := GormDB.Find(msgRec, id)

	return msgRec, res.Error
}

func QueryUnReadMsgById(id uint) (*sql.Row, error) {

	msg := &entity.MessageSend{
		ID: id,
	}

	res := GormDB.
		Where("sender_id = ?", id).
		Where("is_read = ?", constant.FALSE).
		Find(msg)
	if res.Error == nil {
		return nil, res.Error
	}

	return res.Row(), nil
}

func QueryPost(input request.PostQueryInput) (*sql.Rows, *int64, error) {

	var total *int64

	res := GormDB.Model(entity.Post{}).
		Where("id = ?", input.ID).
		Where("author = ?", input.Author).
		Where("content like '?'", input.Content).
		Where("title like '?'", input.Title).
		Scopes(input.PageInfo.Paginate()).
		Count(total)

	rows, err := res.Rows()

	if err != nil {
		log.Logger.Error("query post error: ", zap.Error(err))
		return nil, nil, err
	}

	return rows, total, nil
}
