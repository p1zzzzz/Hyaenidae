package config

type System struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`               // 环境
	Addr    string `mapstructure:"addr" json:"addr" yaml:"addr"`            // 运行端口
	OssType string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"` // Oss类型

}
