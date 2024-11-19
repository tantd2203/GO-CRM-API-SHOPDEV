package global

import (
	"GO-CRM-API-SHOPDEV/pkg/logger"
	"GO-CRM-API-SHOPDEV/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
