package config

import "fmt"

type MySQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	DB           string `mapstructure:"db" json:"db" yaml:"db"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	Conf         string `mapstructure:"conf" json:"conf" yaml:"conf"`
}

func (c *MySQL) ConnectConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.User, c.Password, c.Host, c.Port, c.DB, c.Conf)
}
