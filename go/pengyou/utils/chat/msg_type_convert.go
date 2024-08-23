package chat

import "pengyou/constant"

// Mapping between Redis message types and frontend/backend message types.
var messageTypesMapping = map[int]int{
	constant.MessageTypeText:          constant.TextMessage,
	constant.MessageTypeImage:         constant.UnknownMessageType, // No direct mapping for image messages.
	constant.MessageTypeFile:          constant.UnknownMessageType, // No direct mapping for file messages.
	constant.MessageTypeFileRequest:   constant.FileRequestMessage,
	constant.MessageTypeDisconnect:    constant.UnknownMessageType, // No direct mapping for disconnect messages.
	constant.MessageTypeCutChat:       constant.CutChatMessage,
	constant.MessageTypeEstablishChat: constant.EstablishChatMessage,
}

// Mapping from frontend/backend message types to Redis message types.
var frontendToRedisMapping = map[int]int{
	constant.UnknownMessageType:   0,
	constant.TextMessage:          constant.MessageTypeText,
	constant.FileRequestMessage:   constant.MessageTypeFileRequest,
	constant.FriendRequestMessage: constant.UnknownMessageType,
	constant.EstablishChatMessage: constant.MessageTypeEstablishChat,
	constant.CutChatMessage:       constant.MessageTypeCutChat,
}

// ConvertRdsMsgType converts Redis message types to frontend/backend message types.
func ConvertRdsMsgType(redisMessageType int) int {
	if mappedType, ok := messageTypesMapping[redisMessageType]; ok {
		return mappedType
	}
	return constant.UnknownMessageType
}

// ConvertFrontendMsgType converts frontend/backend message types to Redis message types.
func ConvertFrontendMsgType(frontendMessageType int) int {
	if mappedType, ok := frontendToRedisMapping[frontendMessageType]; ok {
		return mappedType
	}
	return 0 // Default to 0 if no mapping is found.
}
