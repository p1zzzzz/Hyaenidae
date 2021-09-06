package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"Hyaenidae/config"
)

var (
	Hyaenidae_DB                  *gorm.DB
	Hyaenidae_CONFIG              config.Server
	Hyaenidae_VP                  *viper.Viper
	Hyaenidae_LOG                 *zap.Logger
	Hyaenidae_Concurrency_Control = &singleflight.Group{}
)
