package global

import (
	"GO-CRM-API-SHOPDEV/pkg/logger"
	"GO-CRM-API-SHOPDEV/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
