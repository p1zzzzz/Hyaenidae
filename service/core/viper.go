package core

import (
	"Hyaenidae/global"
	"Hyaenidae/utils"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//@function: Viper
//@description: Viper相关配置
//@return: *viper.Viper

func Viper(path ...string) *viper.Viper {
	var config string
	//确定config文件
	if len(path) == 0 {
		if config == "" { // 优先级: 函数值 > 默认值
			config = utils.ConfigFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
		} else {
			config = path[0]
			fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Hyaenidae_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Hyaenidae_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
