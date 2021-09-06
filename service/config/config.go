package config

type Server struct {
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Scrypt    Scrypt    `mapstructure:"scrypt" json:"scrypt" yaml:"scrypt"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	Qiniu     Qiniu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
