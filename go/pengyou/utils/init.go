package utils

import (
	"pengyou/global/config"
	"pengyou/utils/common"
)

func Init(config *config.Config) {
	common.InitSnowflakeIdGenerator(config.Snowflake)
}
