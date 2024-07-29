package config

type ServerConfig struct {
	ServerName   string       `mapstructure:"name"`
	Port         int          `mapstructure:"port"`
	IP           string       `mapstructure:"ip"`
	MysqlConfig  MysqlConfig  `mapstructure:"mysql"`
	LogConfig    LogConfig    `mapstructure:"log"`
	ConsulConfig ConsulConfig `mapstructure:"consul"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
	Database string `mapstructure:"database"`
}

type LogConfig struct {
	Path string `map-structure:"path"`
}

type ConsulConfig struct {
	ConsulHost string   `mapstructure:"consul_host"`
	ConsulPort int      `mapstructure:"consul_port"`
	ConsulTags []string `mapstructure:"tags"`
}
