package constant

import "time"

const (
	REDIS_KEY_USER_CAPTCHA        = "user:captcha:"
	REDIS_KEY_USER_CAPTCHA_EXPIRE = 60 * 5 * time.Second
)
