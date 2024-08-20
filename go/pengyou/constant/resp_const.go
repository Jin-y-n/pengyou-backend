package constant

const (
	SERVER_ERROR = "服务器错误"

	ESTABLISH_WEBSOCKET_CONNECT_FAIL        = "建立websocket连接失败"
	CONNECTED_USER_NOT_FOUND                = "连接用户不存在"
	CONNECTED_USER_NOT_ONLINE               = "连接用户不在线"
	CONNECT_CUTTED                          = "连接断开"
	REDIS_USER_HEARTBEAT_PREFIX             = "user:heartbeat:"
	CONNECTED_USER_SUCCESS                  = "连接用户成功"
	CHAT_ESTABLISH_SUCCESS                  = "聊天建立成功"
	CHAT_ESTABLISH_SUCCESS_FROM             = "聊天建立成功:"
	CHAT_ESTABLISH_FAIL_FROM                = "聊天建立失败:"
	CUT_CHAT_MESSAGE_RESPONSE_SUCCESS       = "聊天断开成功"
	RESP_CHAT_CUTTED_FROM                   = "聊天断开:"
	RESP_CHATTRT_DISCONNECTED               = "聊天用户连接断开"
	RESP_ESTABLISH_CHAT_MESSAGE_FROM_PREFIX = "establish_chat_from:"
	CONNECTED_USER_NOT_FOUND_IN_CHAT_LIST   = "连接用户不在聊天列表中"
	CONNECTED_USER_ALREADY_ONLINE           = "连接用户已经在线"
	CONNECTED_USER_ALREADY_IN_CHAT_LIST     = "连接用户已经在聊天列表中"
	CONNECTED_USER_ALREADY_EXIST            = "连接用户已经存在"

	RESP_DISCONNECT_MESSAGE_PREFIX = "disconnect-user:"
)
