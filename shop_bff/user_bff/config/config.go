package config

type ServerConfig struct {
	ServerName   string       `mapstructure:"name"`
	Port         int          `mapstructure:"port"`
	IP           string       `mapstructure:"ip"`
	JwtConfig    JwtConfig    `mapstructure:"jwt"`
	LogConfig    LogConfig    `mapstructure:"log"`
	ConsulConfig ConsulConfig `mapstructure:"consul"`
}

type JwtConfig struct {
	Key                string `mapstructure:"key"`
	AccessTokenExpire  int64  `mapstructure:"access_token_expire"`
	RefreshTokenExpire int64  `mapstructure:"refresh_token_expire"`
}

type LogConfig struct {
	Path string `map-structure:"path"`
}

type ConsulConfig struct {
	ConsulHost  string   `mapstructure:"consul_host"`
	ConsulPort  int      `mapstructure:"consul_port"`
	ConsulTags  []string `mapstructure:"tags"`
	UserSrvName string   `mapstructure:"user_srv_name"`
}
