package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"temp/config"
	"github.com/go-redis/redis"
)

var (
	GVA_DB *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
	//GVA_Timer timer.Timer = timer.NewTimerTask()
)
