package initalize

import (
	"GO-CRM-API-SHOPDEV/global"
	"GO-CRM-API-SHOPDEV/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
