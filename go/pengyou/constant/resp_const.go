package constant

const (
	ServerError = "服务器错误"

	EstablishWebsocketConnectFail      = "建立websocket连接失败"
	ConnectedUserNotFound              = "连接用户不存在"
	ConnectedUserNotOnline             = "连接用户不在线"
	ConnectCut                         = "连接断开"
	RedisUserHeartbeatPrefix           = "user:heartbeat:"
	ConnectedUserSuccess               = "连接用户成功"
	ChatEstablishSuccess               = "聊天建立成功"
	ChatEstablishSuccessFrom           = "聊天建立成功:"
	ChatEstablishFailFrom              = "聊天建立失败:"
	ChatEstablishFailTo                = "对象聊天建立失败"
	CutChatMessageResponseSuccess      = "聊天断开成功"
	RespChatCutFrom                    = "聊天断开:"
	RespNotConnect                     = "未连接"
	RespChatterDisconnected            = "聊天用户连接断开"
	RespCutChatFailed                  = "聊天断开失败"
	RespEstablishChatMessageFromPrefix = "establish_chat_from:"
	ConnectedUserNotFoundInChatList    = "连接用户不在聊天列表中"
	ConnectedUserAlreadyOnline         = "连接用户已经在线"
	ConnectedUserAlreadyInChatList     = "连接用户已经在聊天列表中"
	ConnectedUserAlreadyExist          = "连接用户已经存在"

	RESP_DISCONNECT_MESSAGE_PREFIX = "disconnect-user:"
)
