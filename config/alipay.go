package config

type Alipay struct {
	AppID        string `mapstructure:"appid" json:"appid" yaml:"appid"`
	PrivateKey   string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
	AliPublicKey string `mapstructure:"ali-public-key" json:"ali-public-key" yaml:"ali-public-key"`
}
