package config

type Scrypt struct {
	Salt string `mapstructure:"salt" json:"salt" yaml:"salt"` // 验证码长度
}
