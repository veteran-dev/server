package config

type Alipay struct {
	AppID            string `mapstructure:"appid" json:"appid" yaml:"appid"`
	PrivateKey       string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
	AppPublicCert    string `mapstructure:"app-public-cert" json:"app-public-cert" yaml:"app-public-cert"`
	AlipayPublicCert string `mapstructure:"alipay-public-cert" json:"alipay-public-cert" yaml:"alipay-public-cert"`
	AlipayRootCert   string `mapstructure:"alipay-root-cert" json:"alipay-root-cert" yaml:"alipay-root-cert"`
}
