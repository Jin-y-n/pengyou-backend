package config

type Email struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Sender   string `mapstructure:"sender" json:"sender" yaml:"sender"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
