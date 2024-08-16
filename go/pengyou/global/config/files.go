package config

type Files struct {
	FilePath     string `mapstructure:"file-path" json:"file-path" yaml:"file-path"`
	BufSize      int    `mapstructure:"buf-size" json:"buf-size" yaml:"buf-size"`
	ReadBufSize  int    `mapstructure:"read-buf-size" json:"read-buf-size" yaml:"read-buf-size"`
	WriteBufSize int    `mapstructure:"write-buf-size" json:"write-buf-size" yaml:"write-buf-size"`

	MesToDBThreshold int64 `mapstructure:"mes-to-db-threshold" json:"mes-to-db-threshold" yaml:"mes-to-db-threshold"`
}
