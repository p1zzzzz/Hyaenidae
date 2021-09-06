package main

import (
	"Hyaenidae/core"
	"Hyaenidae/global"
	"Hyaenidae/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.Hyaenidae_VP = core.Viper()           // 初始化Viper
	global.Hyaenidae_LOG = core.Zap()            // 初始化zap日志库
	global.Hyaenidae_DB = initialize.GormMysql() // gorm连接数据库
	//initialize.Timer()

	if global.Hyaenidae_DB != nil {
		initialize.MysqlTables(global.Hyaenidae_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.Hyaenidae_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
