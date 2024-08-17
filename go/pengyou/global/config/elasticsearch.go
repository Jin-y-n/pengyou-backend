package config

type Elasticsearch struct {
	Host     string   `yaml:"host"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Port     int      `yaml:"port"`
	Nodes    []string `yaml:"nodes"`
	Index    string   `yaml:"index"`
}
