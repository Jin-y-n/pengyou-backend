package constant

import "time"

const (
	REDIS_KEY_USER_CAPTCHA        = "user:captcha:"
	REDIS_KEY_USER_CAPTCHA_EXPIRE = 60 * 5 * time.Second

	REDIS_USER_CHAT_LIST_PREFIX       = "chatting with user"
	REDIS_USER_NODE_PREFIX            = "user_node-"
	REDIS_USER_HEARTBEAT_PREFIX       = "user_heartbeat-"
	REDIS_ESTABLISH_CONNECTION_PREFIX = "establish_connection-"
)
