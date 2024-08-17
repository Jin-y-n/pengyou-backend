package config

type Config struct {
	App           App             `yaml:"app"`
	Email         Email           `yaml:"email"`
	JWT           JWT             `yaml:"jwt"`
	Server        Server          `yaml:"server"`
	Redis         Redis           `yaml:"redis"`
	Zap           Zap             `yaml:"zap"`
	MySQL         MySQL           `yaml:"mysql"`
	Files         Files           `yaml:"files"`
	Elasticsearch []Elasticsearch `yaml:"elasticsearch"`
}

var Cfg *Config
