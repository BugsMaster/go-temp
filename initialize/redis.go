package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"temp/global"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GVA_LOG.Info("redis 连接成功 ！返回值:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
