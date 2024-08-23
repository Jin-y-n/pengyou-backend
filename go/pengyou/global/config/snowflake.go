package config

type Snowflake struct {
	WorkerID uint16 `mapstructure:"worker-id" json:"worker-id" yaml:"worker-id"`
}
