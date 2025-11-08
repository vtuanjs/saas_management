//go:build wireinject

package wire

import (
	"github.com/google/wire"
	system "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/core"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/libs/redis"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/libs/viper"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/libs/zap"
)

func InitConfig() *system.Config {
	wire.Build(
		viper.NewConfig,
	)
	return nil
}

func InitLogger() system.Logger {
	wire.Build(
		viper.NewConfig,
		zap.NewLogger,
	)
	return nil
}

func InitDistributedCache() system.Cache {
	wire.Build(
		viper.NewConfig,
		redis.NewRedisCache,
	)
	return nil
}
