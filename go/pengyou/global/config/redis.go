package config

type Redis struct {
	Username     string   `mapstructure:"username" json:"username" yaml:"username"`
	Addr         string   `mapstructure:"addr" json:"addr" yaml:"addr"`                            // Server address:port
	Password     string   `mapstructure:"password" json:"password" yaml:"password"`                // Password
	DB           int      `mapstructure:"db" json:"db" yaml:"db"`                                  // Database index in single instance mode
	UseCluster   bool     `mapstructure:"use-cluster" json:"use-cluster" yaml:"use-cluster"`       // Whether to use cluster mode
	ClusterAddrs []string `mapstructure:"cluster-addrs" json:"cluster-addrs" yaml:"cluster-addrs"` // List of node addresses in cluster mode
	PoolSize     int      `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"`             // Connection pool size
}
