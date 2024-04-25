package config

type System struct {
	Env           string   `mapstructure:"env" json:"env" yaml:"env"`                // 环境值
	DbType        string   `mapstructure:"db-type" json:"db-type" yaml:"db-type"`    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string   `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix  string   `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr          int      `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP  int      `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int      `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool     `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseRedis      bool     `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	UseMongo      bool     `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`                // 使用mongo
	DituKey       string   `mapstructure:"ditu-key" json:"ditu-key" yaml:"ditu-key"`                   // 环境值
	ChannelCode   string   `mapstructure:"channel-code" json:"channel-code" yaml:"channel-code"`       // 环境值
	POI           []string `mapstructure:"poi" json:"poi" yaml:"poi"`                                  // 环境值
	Reason        []string `mapstructure:"reason" json:"reason" yaml:"reason"`                         // 环境值
}
